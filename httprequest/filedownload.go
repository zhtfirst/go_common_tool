package httprequest

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

// DownloadFile downloads a file from the given URL and saves it to the specified path.
// savePath is the path to save the file, including the file name.
func DownloadFile(ctx context.Context, url string, savePath string) error {
	// Create HTTP client
	response, err := newService("file_download").Do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status code: %d", response.StatusCode)
	}

	// Ensure directory exists
	dir := filepath.Dir(savePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Create file
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write response body to file
	if _, err := io.Copy(file, response.Body); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// ChunkDownload downloads a file in chunks and saves it to the specified path.
func ChunkDownload(ctx context.Context, url string, savePath string, chunkSize int64) error {
	// Create HTTP client
	client := newService("chunk_download")

	totalSize, err := getFileSizeFallback(url)
	if err != nil {
		return fmt.Errorf("failed to get file size: %w", err)
	}

	if totalSize <= 0 {
		return fmt.Errorf("invalid content length: %d", totalSize)
	}

	// Calculate number of chunks
	chunks := totalSize / chunkSize
	if totalSize%chunkSize != 0 {
		chunks++
	}

	// Create temporary directory
	tmpDir := savePath + "_tmp"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create error channel and wait group
	errChan := make(chan error, chunks)
	var wg sync.WaitGroup

	// Download chunks concurrently
	for i := int64(0); i < chunks; i++ {
		wg.Add(1)
		go func(chunkIndex int64) {
			defer wg.Done()

			start := chunkIndex * chunkSize
			end := start + chunkSize - 1
			if end >= totalSize {
				end = totalSize - 1
			}

			// Set Range header
			headers := map[string]string{
				"Range": fmt.Sprintf("bytes=%d-%d", start, end),
			}

			// Download chunk
			resp, err := client.Do(ctx, http.MethodGet, url, headers)
			if err != nil {
				errChan <- fmt.Errorf("failed to download chunk %d: %w", chunkIndex, err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
				errChan <- fmt.Errorf("failed to download chunk %d with status code: %d", chunkIndex, resp.StatusCode)
				return
			}

			// Save chunk
			chunkFile := filepath.Join(tmpDir, fmt.Sprintf("chunk_%d", chunkIndex))
			file, err := os.Create(chunkFile)
			if err != nil {
				errChan <- fmt.Errorf("failed to create chunk file %d: %w", chunkIndex, err)
				return
			}

			_, err = io.Copy(file, resp.Body)
			file.Close()

			if err != nil {
				errChan <- fmt.Errorf("failed to write chunk %d: %w", chunkIndex, err)
				return
			}
		}(i)
	}

	// Wait for all downloads to complete
	wg.Wait()
	close(errChan)

	// Check for errors
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	// Merge chunks
	outFile, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outFile.Close()

	for i := int64(0); i < chunks; i++ {
		chunkFile := filepath.Join(tmpDir, fmt.Sprintf("chunk_%d", i))
		inFile, err := os.Open(chunkFile)
		if err != nil {
			return fmt.Errorf("failed to open chunk %d: %w", i, err)
		}

		_, err = io.Copy(outFile, inFile)
		inFile.Close()

		if err != nil {
			return fmt.Errorf("failed to merge chunk %d: %w", i, err)
		}
	}

	return nil
}

// getFileSizeFallback attempts to get file size using multiple methods in order of efficiency
func getFileSizeFallback(url string) (int64, error) {
	// Try HEAD method first (most efficient)
	if size, err := getFileSize(url); err == nil {
		return size, nil
	}

	// Try Range method second
	if size, err := getFileSizeWithRange(url); err == nil {
		return size, nil
	}

	// Fallback to GET method (least efficient)
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.ContentLength != -1 {
		return resp.ContentLength, nil
	}

	return 0, fmt.Errorf("unable to determine file size using any method")
}

// getFileSize attempts to get file size using HEAD request
func getFileSize(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, fmt.Errorf("HEAD request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HEAD request failed with status: %d", resp.StatusCode)
	}

	if resp.ContentLength == -1 {
		return 0, fmt.Errorf("server did not provide Content-Length")
	}

	return resp.ContentLength, nil
}

// getFileSizeWithRange attempts to get file size using Range request
func getFileSizeWithRange(url string) (int64, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Range", "bytes=0-1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("Range request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		return 0, fmt.Errorf("server does not support Range requests")
	}

	contentRange := resp.Header.Get("Content-Range")
	if contentRange == "" {
		return 0, fmt.Errorf("server did not provide Content-Range header")
	}

	var totalSize int64
	if _, err := fmt.Sscanf(contentRange, "bytes 0-1/%d", &totalSize); err != nil {
		return 0, fmt.Errorf("failed to parse Content-Range: %w", err)
	}

	return totalSize, nil
}
