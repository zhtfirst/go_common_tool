package tool

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/spf13/cast"
)

// DayDiff 计算两个日期之间的天数差
func DayDiff(beginDate, endDate string) (days int) {
	begin, _ := time.Parse("2006-01-02", beginDate)
	end, _ := time.Parse("2006-01-02", endDate)
	days = cast.ToInt(end.Sub(begin).Hours() / 24)
	return
}

// DateFormat 日期格式化 yyyy-mm-dd
func DateFormat(date string) string {
	dateTime, _ := time.Parse(time.RFC3339, date)
	return dateTime.Format("2006-01-02")
}

// GetAddDayTime 获取几年、几月、几天后的时间
func GetAddDayTime(year, month, day int) string {
	return time.Now().AddDate(year, month, day).Format("2006-01-02 15:04:05")
}

// GetNowTime 获取当前时间
func GetNowTime() string {
	return cast.ToString(time.Now().Format("2006-01-02 15:04:05"))
}

// FormatTimeToString 将时间类型转换为 "2006-01-02 15:04" 格式的字符串
// 支持 time.Time 和 sql.NullTime 两种输入类型
func FormatTimeToString(timeObj interface{}) (string, error) {
	var t time.Time
	var valid bool

	switch v := timeObj.(type) {
	case time.Time:
		t = v
		valid = true
	case sql.NullTime:
		if !v.Valid {
			return "", fmt.Errorf("时间值为空")
		}
		t = v.Time
		valid = true
	default:
		return "", fmt.Errorf("不支持的时间类型: %T", timeObj)
	}

	if !valid {
		return "", fmt.Errorf("无效的时间值")
	}

	// 定义输出格式 (去掉秒和时区信息)
	const layout = "2006-01-02 15:04"
	return t.Format(layout), nil
}

// RemoveTimezone 从时间字符串中移除时区信息
func RemoveTimezone(ctx context.Context, timeStr string) string {
	// 解析带时区的时间字符串
	t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", timeStr)
	if err != nil {
		return ""
	}

	// 格式化为不带时区的时间字符串
	return t.Format("2006-01-02 15:04:05")
}
