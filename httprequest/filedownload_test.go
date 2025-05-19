package httprequest

import (
	"context"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	ctx := context.Background()
	// err := DownloadFile(ctx, "https://www.baidu.com", "file.txt")
	err := DownloadFile(ctx, "https://www.httpwatch.com/httpgallery/chunked/chunkedimage.aspx", "./test/file.jpg")
	if err != nil {
		t.Fatalf("DownloadFile failed: %v", err)
	}
}

func TestChunkDownload(t *testing.T) {
	ctx := context.Background()
	// err := ChunkDownload(ctx, "https://www.httpwatch.com/httpgallery/chunked/chunkedimage.aspx", "file.jpg", 1024)
	err := ChunkDownload(ctx, "https://sample-videos.com/video123/mp4/720/big_buck_bunny_720p_1mb.mp4", "file.jpg", 1024)
	if err != nil {
		t.Fatalf("ChunkDownload failed: %v", err)
	}
}
