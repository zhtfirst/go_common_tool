package tool

import (
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

// TestGoSafe_Normal 测试正常执行的情况
func TestGoSafe_Normal(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)

	executed := false
	GoSafe(func() {
		executed = true
		wg.Done()
	})

	wg.Wait()
	if !executed {
		t.Error("GoSafe 未能正常执行函数")
	}
}

// TestGoSafe_WithPanic 测试发生 panic 时的恢复能力
func TestGoSafe_WithPanic(t *testing.T) {
	// 清理之前的日志文件
	logDir := "./log"
	os.RemoveAll(logDir)

	var wg sync.WaitGroup
	wg.Add(1)

	GoSafe(func() {
		defer wg.Done()
		// 制造一个 panic
		panic("测试 panic")
	})

	// 等待 goroutine 完成
	wg.Wait()

	// 给一点时间让日志文件写入
	time.Sleep(500 * time.Millisecond)

	// 验证是否创建了崩溃日志
	files, err := filepath.Glob(logDir + "/*.crash")
	if err != nil {
		t.Errorf("无法列出崩溃日志文件: %v", err)
	}

	if len(files) == 0 {
		t.Error("未创建崩溃日志文件")
		return
	}

	// 读取崩溃日志内容
	content, err := os.ReadFile(files[0])
	if err != nil {
		t.Errorf("无法读取崩溃日志文件: %v", err)
		return
	}

	if !strings.Contains(string(content), "测试 panic") {
		t.Errorf("崩溃日志未包含 panic 信息")
	}

	// 清理
	os.RemoveAll(logDir)
}

// TestRecoverPanic 直接测试 RecoverPanic 函数
func TestRecoverPanic(t *testing.T) {
	// 清理之前的日志文件
	logDir := "./log"
	os.RemoveAll(logDir)

	func() {
		defer RecoverPanic()
		panic("直接测试 panic 恢复")
	}()

	// 给一点时间让日志文件写入
	time.Sleep(500 * time.Millisecond)

	// 验证是否创建了崩溃日志
	files, err := filepath.Glob(logDir + "/*.crash")
	if err != nil {
		t.Errorf("无法列出崩溃日志文件: %v", err)
	}

	if len(files) == 0 {
		t.Error("未创建崩溃日志文件")
		return
	}

	// 清理
	os.RemoveAll(logDir)
}
