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

// RangeHeader 用于Range请求的请求头结构体
type RangeHeader struct {
	Range string `header:"Range"`
}

// DownloadFile 从指定URL下载文件并保存到指定路径
// ctx: 上下文，用于控制请求的生命周期
// url: 文件下载地址
// savePath: 保存文件的路径，包括文件名
// 返回值: 错误信息，如果下载成功则返回nil
func DownloadFile(ctx context.Context, url string, savePath string) error {
	// 创建HTTP客户端
	response, err := newService("file_download").Do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("创建HTTP请求失败: %w", err)
	}
	defer response.Body.Close()

	// 检查响应状态
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("下载失败，状态码: %d", response.StatusCode)
	}

	// 确保目录存在
	dir := filepath.Dir(savePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 创建文件
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	defer file.Close()

	// 将响应内容写入文件
	if _, err := io.Copy(file, response.Body); err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	return nil
}

// ChunkDownload 分块下载文件并保存到指定路径(可以下载任何类型的文件)
// ctx: 上下文，用于控制请求的生命周期
// url: 文件下载地址
// savePath: 保存文件的路径，包括文件名
// chunkSize: 每个分块的大小（字节）
// 返回值: 错误信息，如果下载成功则返回nil
func ChunkDownload(ctx context.Context, url string, savePath string, chunkSize int64) error {
	// 创建HTTP客户端
	client := newService("chunk_download")

	// 获取文件大小
	totalSize, err := getFileSizeFallback(url)
	if err != nil {
		// 如果无法获取文件大小，则使用普通下载方式
		return DownloadFile(ctx, url, savePath)
	}

	if totalSize <= 0 {
		return fmt.Errorf("无效的内容长度: %d", totalSize)
	}

	// 计算分块数量
	chunks := totalSize / chunkSize
	if totalSize%chunkSize != 0 {
		chunks++
	}

	// 创建临时目录
	tmpDir := savePath + "_tmp"
	if err := os.MkdirAll(tmpDir, 0755); err != nil {
		return fmt.Errorf("创建临时目录失败: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建错误通道和等待组
	errChan := make(chan error, chunks)
	var wg sync.WaitGroup

	// 并发下载分块
	for i := int64(0); i < chunks; i++ {
		wg.Add(1)
		go func(chunkIndex int64) {
			defer wg.Done()

			// 计算当前分块的起始和结束位置
			start := chunkIndex * chunkSize
			end := start + chunkSize - 1
			if end >= totalSize {
				end = totalSize - 1
			}

			// 设置Range请求头
			rangeHeader := RangeHeader{
				Range: fmt.Sprintf("bytes=%d-%d", start, end),
			}

			// 下载分块
			resp, err := client.Do(ctx, http.MethodGet, url, rangeHeader)
			if err != nil {
				errChan <- fmt.Errorf("下载分块%d失败: %w", chunkIndex, err)
				return
			}
			defer resp.Body.Close()

			// 检查响应状态
			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusPartialContent {
				errChan <- fmt.Errorf("下载分块%d失败，状态码: %d", chunkIndex, resp.StatusCode)
				return
			}

			// 保存分块
			chunkFile := filepath.Join(tmpDir, fmt.Sprintf("chunk_%d", chunkIndex))
			file, err := os.Create(chunkFile)
			if err != nil {
				errChan <- fmt.Errorf("创建分块文件%d失败: %w", chunkIndex, err)
				return
			}

			// 将分块内容写入文件
			_, err = io.Copy(file, resp.Body)
			file.Close()

			if err != nil {
				errChan <- fmt.Errorf("写入分块%d失败: %w", chunkIndex, err)
				return
			}
		}(i)
	}

	// 等待所有下载完成
	wg.Wait()
	close(errChan)

	// 检查是否有错误
	for err := range errChan {
		if err != nil {
			return err
		}
	}

	// 合并分块
	outFile, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("创建输出文件失败: %w", err)
	}
	defer outFile.Close()

	// 按顺序合并所有分块
	for i := int64(0); i < chunks; i++ {
		chunkFile := filepath.Join(tmpDir, fmt.Sprintf("chunk_%d", i))
		inFile, err := os.Open(chunkFile)
		if err != nil {
			return fmt.Errorf("打开分块%d失败: %w", i, err)
		}

		// 将分块内容写入最终文件
		_, err = io.Copy(outFile, inFile)
		inFile.Close()

		if err != nil {
			return fmt.Errorf("合并分块%d失败: %w", i, err)
		}
	}

	return nil
}

// getFileSizeFallback 尝试使用多种方法获取文件大小，按效率顺序尝试
// url: 文件URL
// 返回值: 文件大小和错误信息
func getFileSizeFallback(url string) (int64, error) {
	// 首先尝试HEAD方法（最高效）
	if size, err := getFileSize(url); err == nil {
		return size, nil
	}

	// 其次尝试Range方法
	if size, err := getFileSizeWithRange(url); err == nil {
		return size, nil
	}

	// 最后尝试GET方法（效率最低）
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("GET请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.ContentLength != -1 {
		return resp.ContentLength, nil
	}

	return 0, fmt.Errorf("无法通过任何方法确定文件大小")
}

// getFileSize 尝试使用HEAD请求获取文件大小
// url: 文件URL
// 返回值: 文件大小和错误信息
func getFileSize(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, fmt.Errorf("HEAD请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HEAD请求失败，状态码: %d", resp.StatusCode)
	}

	if resp.ContentLength == -1 {
		return 0, fmt.Errorf("服务器未提供Content-Length")
	}

	return resp.ContentLength, nil
}

// getFileSizeWithRange 尝试使用Range请求获取文件大小
// url: 文件URL
// 返回值: 文件大小和错误信息
func getFileSizeWithRange(url string) (int64, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("创建请求失败: %w", err)
	}

	req.Header.Set("Range", "bytes=0-1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("Range请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent {
		return 0, fmt.Errorf("服务器不支持Range请求")
	}

	contentRange := resp.Header.Get("Content-Range")
	if contentRange == "" {
		return 0, fmt.Errorf("服务器未提供Content-Range头")
	}

	var totalSize int64
	if _, err := fmt.Sscanf(contentRange, "bytes 0-1/%d", &totalSize); err != nil {
		return 0, fmt.Errorf("解析Content-Range失败: %w", err)
	}

	return totalSize, nil
}
