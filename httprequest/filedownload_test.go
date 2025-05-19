/*
 * @Author: gavin
 * @Date: 2025-05-19 14:58:24
 * @LastEditors: gavin
 * @LastEditTime: 2025-05-19 15:48:14
 * @FilePath: /httprequest/filedownload_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package httprequest

import (
	"context"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	ctx := context.Background()
	err := DownloadFile(ctx, "https://www.httpwatch.com/httpgallery/chunked/chunkedimage.aspx", "./test/download/file.jpg")
	if err != nil {
		t.Fatalf("DownloadFile failed: %v", err)
	}
}

func TestChunkDownload(t *testing.T) {
	ctx := context.Background()
	//err := ChunkDownload(ctx, "https://www.httpwatch.com/httpgallery/chunked/chunkedimage.aspx", "./test/chunkdownload/file.jpg", 1024)
	err := ChunkDownload(ctx, "https://vdept3.bdstatic.com/mda-ra3ivn98qcw5niim/sc/cae_h264/1735998903757667359/mda-ra3ivn98qcw5niim.mp4?v_from_s=hkapp-haokan-hbf&auth_key=1747650790-0-0-7a4a1b796802a9092d12beb4f91a58e2&bcevod_channel=searchbox_feed&pd=1&cr=2&cd=0&pt=3&logid=1990645079&vid=13292202250349129751&klogid=1990645079&abtest=", "./test/chunkdownload/vedio.mp4", 1024)
	if err != nil {
		t.Fatalf("ChunkDownload failed: %v", err)
	}
}

// TestDownloadVideoAndZip 测试下载MP4视频和ZIP压缩包
func TestDownloadVideoAndZip(t *testing.T) {
	ctx := context.Background()

	// 测试下载MP4视频
	// 使用一个公开的、稳定的MP4视频URL
	videoURL := "https://filesamples.com/samples/video/mp4/sample_960x540.mp4"
	videoPath := "./test/video/sample_video.mp4"
	t.Logf("开始下载MP4视频: %s", videoURL)
	err := ChunkDownload(ctx, videoURL, videoPath, 1024*100) // 使用100KB的块大小
	if err != nil {
		t.Fatalf("下载MP4视频失败: %v", err)
	}
	t.Logf("MP4视频下载成功: %s", videoPath)

	// 测试下载ZIP压缩包
	// 使用一个公开的、稳定的ZIP文件URL
	zipURL := "https://www.learningcontainer.com/wp-content/uploads/2020/05/sample-zip-file.zip"
	zipPath := "./test/zip/sample.zip"
	t.Logf("开始下载ZIP压缩包: %s", zipURL)
	err = ChunkDownload(ctx, zipURL, zipPath, 1024*500) // 使用500KB的块大小
	if err != nil {
		t.Fatalf("下载ZIP压缩包失败: %v", err)
	}
	t.Logf("ZIP压缩包下载成功: %s", zipPath)
}
