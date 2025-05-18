package tool

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gfile"
	"time"
)

// GoSafe 启动一个带有panic恢复机制的goroutine
// 使用示例：
//
//	GoSafe(func() {
//	    ProcessData(data)
//	})
func GoSafe(fn func()) {
	go func() {
		defer RecoverPanic()
		fn()
	}()
}

// RecoverPanic 恢复panic并保存崩溃日志
func RecoverPanic() {
	if r := recover(); r != nil {
		panicMsg := fmt.Sprintf("%v", r)
		err := gerror.NewSkip(1, panicMsg)

		// 确保日志目录存在
		logDir := "./log"
		if err := gfile.Mkdir(logDir); err != nil {
			fmt.Printf("创建日志目录失败: %v\n", err)
			return
		}

		// 使用更可读的时间格式
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		crashFileName := fmt.Sprintf("%s/%s.crash", logDir, timestamp)

		// 写入文件并处理可能的错误
		if err := gfile.PutContents(crashFileName, gerror.Stack(err)); err != nil {
			fmt.Printf("写入崩溃日志失败: %v\n", err)
		}
	}
}
