package tool

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cast"
)

// DayDiff 计算两个日期字符串之间的天数差（格式：2006-01-02）
// beginDate: 起始日期字符串
// endDate: 结束日期字符串
// 返回值: 天数差（int），如果格式错误返回0
func DayDiff(beginDate, endDate string) (days int) {
	begin, _ := time.Parse("2006-01-02", beginDate)
	end, _ := time.Parse("2006-01-02", endDate)
	days = cast.ToInt(end.Sub(begin).Hours() / 24)
	return
}

// DateFormat 将RFC3339格式的时间字符串转为yyyy-mm-dd格式
// date: RFC3339格式时间字符串
// 返回值: yyyy-mm-dd格式字符串，格式错误返回空字符串
func DateFormat(date string) string {
	dateTime, _ := time.Parse(time.RFC3339, date)
	return dateTime.Format("2006-01-02")
}

// GetAddDayTime 获取当前时间加指定年、月、天后的时间字符串（格式：yyyy-mm-dd HH:MM:SS）
// year: 增加的年数
// month: 增加的月数
// day: 增加的天数
// 返回值: 新时间字符串
func GetAddDayTime(year, month, day int) string {
	return time.Now().AddDate(year, month, day).Format("2006-01-02 15:04:05")
}

// GetNowTime 获取当前时间字符串（格式：yyyy-mm-dd HH:MM:SS）
// 返回值: 当前时间字符串
func GetNowTime() string {
	return cast.ToString(time.Now().Format("2006-01-02 15:04:05"))
}

// FormatTimeToString 将时间类型转换为 "2006-01-02 15:04" 格式的字符串
// 支持 time.Time 和 sql.NullTime 两种输入类型
// timeObj: 时间对象（time.Time 或 sql.NullTime）
// 返回值: 格式化字符串，或错误
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

// RemoveTimezone 从时间字符串中移除时区信息，返回不带时区的时间字符串
// ctx: 上下文参数（可忽略）
// timeStr: 带时区的时间字符串（格式：2006-01-02 15:04:05 -0700 MST）
// 返回值: 不带时区的时间字符串，格式错误返回空字符串
func RemoveTimezone(ctx context.Context, timeStr string) string {
	// 解析带时区的时间字符串
	t, err := time.Parse("2006-01-02 15:04:05 -0700 MST", timeStr)
	if err != nil {
		return ""
	}

	// 格式化为不带时区的时间字符串
	return t.Format("2006-01-02 15:04:05")
}

// GetCurTimeStr 获取当前日期时间字符串（格式：2006-01-02 15:04:05）
// 返回值: 当前时间字符串
func GetCurTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetCurDateStr 获取当前日期字符串（格式：20060102）
// 返回值: 当前日期字符串
func GetCurDateStr() string {
	return time.Now().Format("20060102")
}

// GetTomorrowDateStr 获取明天日期字符串（格式：20060102）
// 返回值: 明天日期字符串
func GetTomorrowDateStr() string {
	return time.Now().AddDate(0, 0, 1).Format("20060102")
}

// GetBeforeDateStr 获取指定天数后的日期字符串（格式：20060102）
// count: 天数偏移，0为今天，负数为过去，正数为未来
// 返回值: 日期字符串
func GetBeforeDateStr(count int) string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, count)
	yes_day := yesTime.Format("20060102")

	return yes_day
}

// GetCurDayStr 获取当前日期字符串（格式：2006-01-02）
// 返回值: 当前日期字符串
func GetCurDayStr() string {
	return time.Now().Format("2006-01-02")
}

// GetYesterDayStr 获取昨天的日期字符串（格式：2006-01-02）
// 返回值: 昨天日期字符串
func GetYesterDayStr() string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	yes_day := yesTime.Format("2006-01-02")

	return yes_day
}

// GetMonthDayStr 时间戳转化成 MM-DD 格式字符串
// timestamp: 秒级时间戳
// 返回值: MM-DD 格式字符串
func GetMonthDayStr(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("01-02")
}

// GetAddDayStr 获取指定天数后的日期字符串（格式：2006-01-02 00:00:00）
// days: 天数偏移
// 返回值: 日期字符串
func GetAddDayStr(days int) string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, days)
	retDay := yesTime.Format("2006-01-02 00:00:00")

	return retDay
}

// GetNextWeekStart 获取下周一的日期字符串（格式：2006-01-02 00:00:00）
// 返回值: 下周一日期字符串
func GetNextWeekStart() string {
	nTime := time.Now()
	weekDay := int(nTime.Weekday())
	return GetAddDayStr(7 - weekDay + 1)
}

// GetNextMonthStart 获取下月一号的日期字符串（格式：2006-01-02 00:00:00）
// 返回值: 下月一号日期字符串
func GetNextMonthStart() string {
	nTime := time.Now()
	year := nTime.Year()
	month := nTime.Month()

	month = month + 1
	if month > 12 {
		month = 1
		year = year + 1
	}

	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02 00:00:00")
}

// GetBeforeDayStr 获取指定天数前的日期字符串（格式：2006-01-02）
// count: 天数偏移，负数为未来，正数为过去
// 返回值: 日期字符串
func GetBeforeDayStr(count int) string {
	nTime := time.Now()
	tmpTime := nTime.AddDate(0, 0, 0-count)
	beforeDay := tmpTime.Format("2006-01-02")

	return beforeDay
}

// GetTimestamp 获取当前时间的毫秒级时间戳
// 返回值: 毫秒级时间戳（int）
func GetTimestamp() int {
	return int(time.Now().UnixNano() / 1e6)
}

// GetCurYearMonth 获取当前年月字符串（格式：200601）
// 返回值: 当前年月字符串
func GetCurYearMonth() string {
	return time.Now().Format("200601")
}

// GetNextYearMonth 获取下月年月字符串（格式：200601）
// 返回值: 下月年月字符串
func GetNextYearMonth() string {
	return time.Now().AddDate(0, 1, 0).Format("200601")
}

// GetDateValue 获取当前日期的int值（格式：yyyyMMdd）
// 返回值: 当前日期int值
func GetDateValue() int {
	v, _ := strconv.ParseInt(time.Now().Format("20060102"), 10, 32)
	return int(v)
}

// GetDayLeftover 获取到明天凌晨剩余的秒数
// 返回值: 剩余秒数
func GetDayLeftover() int {
	now := time.Now()
	cur_time := now.Unix()
	timeStr := now.Format("2006-01-02")

	// 使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	return int(t.Unix() + 1 - cur_time)
}

// GetRemainingTimeLen 获取到指定时间的剩余秒数
// timeStr: 目标时间字符串（格式：2006-01-02 15:04:05）
// 返回值: 剩余秒数，格式错误返回0
func GetRemainingTimeLen(timeStr string) int {
	now := time.Now()
	cur_time := now.Unix()
	// 使用Parse 默认获取为UTC时区 需要获取本地时区 所以使用ParseInLocation
	t, e := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)

	if nil != e {
		return 0
	}

	return int(t.Unix() + 1 - cur_time)
}

// GetTheMonthAgoDayStr 获取指定月数前的日期时间字符串（格式：2006-01-02 15:04:05）
// month: 月数偏移
// 返回值: 日期时间字符串
func GetTheMonthAgoDayStr(month int) string {
	nowTime := time.Now()
	getTime := nowTime.AddDate(0, 0-month, 0)    // 年，月，日   获取一个月前的时间
	return getTime.Format("2006-01-02 15:04:05") // 获取的时间的格式
}

// TimeStrToTimestamp 时间字符串转秒级时间戳
// timestr: 时间字符串（格式：2006-01-02 15:04:05）
// 返回值: 秒级时间戳，格式错误返回0
func TimeStrToTimestamp(timestr string) int {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	if err != nil {
		return 0
	}

	return int(t.Unix())
}

// GetCurTimestamp 获取当前的秒级时间戳
// 返回值: 当前秒级时间戳（int）
func GetCurTimestamp() int {
	return int(time.Now().Unix())
}

// GetSubDayCount 获取指定时间到现在的天数
// timestr: 时间字符串（格式：2006-01-02 15:04:05）
// 返回值: 天数差，格式错误返回0
func GetSubDayCount(timestr string) int {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", timestr, time.Local)
	if nil != err {
		return 0
	}
	days := time.Now().Sub(t).Hours() / 24
	return int(days)
}

// TodayLastTime 计算今天剩余时间，单位秒
// 返回值: 剩余秒数
func TodayLastTime() int64 {
	var (
		t             = time.Now()
		TodayZeroTime = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
		TodayEndTime  = TodayZeroTime + 86399
	)
	return TodayEndTime - t.Unix()
}

// TodayPastTime 计算今天已经过去的时间，单位秒
// 返回值: 已过去秒数
func TodayPastTime() int64 {
	var (
		t             = time.Now()
		TodayZeroTime = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	)
	return t.Unix() - TodayZeroTime
}

// WeekLastTime 计算本周剩余时间，单位秒
// 返回值: 剩余秒数
func WeekLastTime() int64 {
	var (
		t            = time.Now()
		WeekZeroTime = time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), 0, 0, 0, 0, t.Location()).Unix()
		WeekEndTime  = WeekZeroTime + 604799
	)
	return WeekEndTime - t.Unix()
}

// WeekStartTime 计算本周开始时间，单位秒
// 返回值: 本周开始时间戳
func WeekStartTime() int64 {
	var (
		t            = time.Now()
		WeekZeroTime = time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), 0, 0, 0, 0, t.Location()).Unix()
	)
	return WeekZeroTime
}

// WeekEndTime 计算本周结束时间，单位秒
// 返回值: 本周结束时间戳
func WeekEndTime() int64 {
	var (
		t            = time.Now()
		WeekZeroTime = time.Date(t.Year(), t.Month(), t.Day()-int(t.Weekday()), 0, 0, 0, 0, t.Location()).Unix()
		endTime      = WeekZeroTime + 604799
	)
	return endTime
}

// GetYMDFormat 获取指定天数偏移的日期字符串（格式：2006-01-02）
// t: 天数偏移，0为今天，-1为昨天
// 返回值: 日期字符串
func GetYMDFormat(t int) string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, t)
	logDay := yesTime.Format("2006-01-02")
	return logDay
}

// TimeStampToTime 时间戳转日期字符串（格式：2006-01-02 15:04:05）
// timestamp: 秒级时间戳
// 返回值: 日期时间字符串
func TimeStampToTime(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}

// GetTimeRange 获取指定天数的时间范围（当天0点到23:59:59）
// count: 天数偏移
// 返回值: (开始时间字符串, 结束时间字符串)
func GetTimeRange(count int) (startTimeFormat string, endTimeFormat string) {
	t := time.Now()
	Time := t.AddDate(0, 0, count)
	// 获取当天开始时间
	start := Time
	startTimeFormat = start.Format("2006-01-02 00:00:00")
	// 获取当天结束时间
	endTime := time.Date(Time.Year(), Time.Month(), Time.Day(), 23, 59, 59, 0, Time.Location()).Unix()
	endTimeFormat = TimeStampToTime(endTime)

	return startTimeFormat, endTimeFormat
}

// GetTodayRemainSeconds 获取今天剩余秒数
// 返回值: 剩余秒数
func GetTodayRemainSeconds() int {
	now := time.Now()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	duration := tomorrow.Sub(now)
	return int(duration.Seconds())
}

// TruncateTimeToStart 获取某天零点的时间戳
// start: 目标时间
// 返回值: 零点时间戳
func TruncateTimeToStart(start time.Time) int64 {
	year, month, day := start.Date()
	startOfDay := time.Date(year, month, day, 0, 0, 0, 0, start.Location())
	return startOfDay.Unix()
}

// TruncateTimeToEnd 获取某天23:59:59的时间戳
// end: 目标时间
// 返回值: 23:59:59时间戳
func TruncateTimeToEnd(end time.Time) int64 {
	year, month, day := end.Date()
	endOfDay := time.Date(year, month, day, 23, 59, 59, 0, end.Location())
	return endOfDay.Unix()
}
