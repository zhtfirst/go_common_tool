/*
 * @Author: gavin v_zhangtao15@tal.com
 * @Date: 2025-05-16 18:04:30
 * @LastEditors: gavin zhtfirst@163.com
 * @LastEditTime: 2025-05-17 09:37:35
 * @FilePath: /go_common_tool/tool/time_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tool

import (
	"context"
	"database/sql"
	"testing"
	"time"
)

func TestDayDiff(t *testing.T) {
	d1 := "2023-01-01"
	d2 := "2023-01-10"
	days := DayDiff(d1, d2)
	if days != 9 {
		t.Errorf("DayDiff(%v, %v) == %v, expected 9", d1, d2, days)
	}
	// 反向测试
	days2 := DayDiff(d2, d1)
	if days2 != -9 {
		t.Errorf("DayDiff(%v, %v) == %v, expected -9", d2, d1, days2)
	}
}

func TestDateFormat(t *testing.T) {
	// RFC3339格式
	formatted := DateFormat("2023-01-01T15:04:05Z")
	if formatted != "2023-01-01" {
		t.Errorf("DateFormat error: got %v", formatted)
	}
	// 非法格式
	if DateFormat("bad") != "" {
		t.Error("DateFormat with bad input should return empty string")
	}
	// 其它RFC3339合法格式
	formatted2 := DateFormat("2023-01-01T00:00:00+08:00")
	if formatted2 != "2023-01-01" {
		t.Errorf("DateFormat error: got %v", formatted2)
	}
}

func TestGetAddDayTime(t *testing.T) {
	future := GetAddDayTime(0, 0, 1)
	expected := time.Now().AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	if future != expected {
		t.Errorf("GetAddDayTime(0, 0, 1) == %v, expected %v", future, expected)
	}
}

func TestGetNowTime(t *testing.T) {
	now := GetNowTime()
	if len(now) != 19 {
		t.Errorf("GetNowTime() length error: %v", now)
	}
}

func TestFormatTimeToString(t *testing.T) {
	t1 := time.Date(2023, 1, 1, 15, 4, 0, 0, time.UTC)
	formatted, err := FormatTimeToString(t1)
	if err != nil || formatted != "2023-01-01 15:04" {
		t.Errorf("FormatTimeToString(t1) == %v, %v", formatted, err)
	}
	nullTime := sql.NullTime{Valid: false}
	_, err = FormatTimeToString(nullTime)
	if err == nil {
		t.Error("FormatTimeToString(null) should error")
	}
	_, err = FormatTimeToString(123)
	if err == nil {
		t.Error("FormatTimeToString(123) should error")
	}
}

func TestRemoveTimezone(t *testing.T) {
	ctx := context.TODO()
	timeStr := "2023-01-01 15:04:05 +0000 UTC"
	result := RemoveTimezone(ctx, timeStr)
	expected := "2023-01-01 15:04:05"
	if result != expected {
		t.Errorf("RemoveTimezone error: %v", result)
	}
	if RemoveTimezone(ctx, "bad") != "" {
		t.Error("RemoveTimezone with bad input should return empty string")
	}
}

func TestGetCurTimeStr(t *testing.T) {
	res := GetCurTimeStr()
	if len(res) != 19 {
		t.Errorf("GetCurTimeStr() length error: %v", res)
	}
}

func TestGetCurDateStr(t *testing.T) {
	res := GetCurDateStr()
	if len(res) != 8 {
		t.Errorf("GetCurDateStr() length error: %v", res)
	}
}

func TestGetTomorrowDateStr(t *testing.T) {
	tomorrow := time.Now().AddDate(0, 0, 1).Format("20060102")
	if GetTomorrowDateStr() != tomorrow {
		t.Errorf("GetTomorrowDateStr() == %v, expected %v", GetTomorrowDateStr(), tomorrow)
	}
}

func TestGetBeforeDateStr(t *testing.T) {
	today := time.Now().Format("20060102")
	if GetBeforeDateStr(0) != today {
		t.Errorf("GetBeforeDateStr(0) == %v, expected %v", GetBeforeDateStr(0), today)
	}
}

func TestGetCurDayStr(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	if GetCurDayStr() != today {
		t.Errorf("GetCurDayStr() == %v, expected %v", GetCurDayStr(), today)
	}
}

func TestGetYesterDayStr(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	if GetYesterDayStr() != yesterday {
		t.Errorf("GetYesterDayStr() == %v, expected %v", GetYesterDayStr(), yesterday)
	}
}

func TestGetMonthDayStr(t *testing.T) {
	ts := time.Date(2023, 6, 11, 0, 0, 0, 0, time.UTC).Unix()
	if GetMonthDayStr(ts) != "06-11" {
		t.Errorf("GetMonthDayStr error: %v", GetMonthDayStr(ts))
	}
}

func TestGetAddDayStr(t *testing.T) {
	res := GetAddDayStr(1)
	if len(res) != 19 {
		t.Errorf("GetAddDayStr(1) length error: %v", res)
	}
}

func TestGetNextWeekStart(t *testing.T) {
	res := GetNextWeekStart()
	if len(res) != 19 {
		t.Errorf("GetNextWeekStart() length error: %v", res)
	}
}

func TestGetNextMonthStart(t *testing.T) {
	res := GetNextMonthStart()
	if len(res) != 19 {
		t.Errorf("GetNextMonthStart() length error: %v", res)
	}
}

func TestGetBeforeDayStr(t *testing.T) {
	res := GetBeforeDayStr(1)
	if len(res) != 10 {
		t.Errorf("GetBeforeDayStr(1) length error: %v", res)
	}
}

func TestGetTimestamp(t *testing.T) {
	ts := GetTimestamp()
	if ts <= 0 {
		t.Errorf("GetTimestamp() error: %v", ts)
	}
}

func TestGetCurYearMonth(t *testing.T) {
	res := GetCurYearMonth()
	if len(res) != 6 {
		t.Errorf("GetCurYearMonth() length error: %v", res)
	}
}

func TestGetNextYearMonth(t *testing.T) {
	res := GetNextYearMonth()
	if len(res) != 6 {
		t.Errorf("GetNextYearMonth() length error: %v", res)
	}
}

func TestGetDateValue(t *testing.T) {
	val := GetDateValue()
	if val <= 20220101 {
		t.Errorf("GetDateValue error: %v", val)
	}
}

func TestGetDayLeftover(t *testing.T) {
	left := GetDayLeftover()
	if left <= 0 || left > 86400 {
		t.Errorf("GetDayLeftover error: %v", left)
	}
}

func TestGetRemainingTimeLen(t *testing.T) {
	now := time.Now().Add(10 * time.Second).Format("2006-01-02 15:04:05")
	left := GetRemainingTimeLen(now)
	if left <= 0 || left > 86410 {
		t.Errorf("GetRemainingTimeLen error: %v", left)
	}
	if GetRemainingTimeLen("bad") != 0 {
		t.Error("GetRemainingTimeLen with bad input should return 0")
	}
}

func TestGetTheMonthAgoDayStr(t *testing.T) {
	res := GetTheMonthAgoDayStr(1)
	if len(res) != 19 {
		t.Errorf("GetTheMonthAgoDayStr(1) length error: %v", res)
	}
}

func TestTimeStrToTimestamp(t *testing.T) {
	ts := TimeStrToTimestamp("2023-01-01 00:00:00")
	if ts == 0 {
		t.Error("TimeStrToTimestamp error")
	}
	if TimeStrToTimestamp("bad") != 0 {
		t.Error("TimeStrToTimestamp with bad input should return 0")
	}
}

func TestGetCurTimestamp(t *testing.T) {
	ts := GetCurTimestamp()
	if ts <= 0 {
		t.Errorf("GetCurTimestamp() error: %v", ts)
	}
}

func TestGetSubDayCount(t *testing.T) {
	ts := time.Now().AddDate(0, 0, -2).Format("2006-01-02 15:04:05")
	if GetSubDayCount(ts) < 1 {
		t.Errorf("GetSubDayCount error: %v", GetSubDayCount(ts))
	}
	if GetSubDayCount("bad") != 0 {
		t.Error("GetSubDayCount with bad input should return 0")
	}
}

func TestTodayLastTime(t *testing.T) {
	left := TodayLastTime()
	if left <= 0 || left > 86400 {
		t.Errorf("TodayLastTime error: %v", left)
	}
}

func TestTodayPastTime(t *testing.T) {
	past := TodayPastTime()
	if past < 0 || past > 86400 {
		t.Errorf("TodayPastTime error: %v", past)
	}
}

func TestWeekLastTime(t *testing.T) {
	left := WeekLastTime()
	if left < 0 || left > 604800 {
		t.Errorf("WeekLastTime error: %v", left)
	}
}

func TestWeekStartTime(t *testing.T) {
	start := WeekStartTime()
	if start <= 0 {
		t.Errorf("WeekStartTime error: %v", start)
	}
}

func TestWeekEndTime(t *testing.T) {
	end := WeekEndTime()
	if end <= 0 {
		t.Errorf("WeekEndTime error: %v", end)
	}
}

func TestGetYMDFormat(t *testing.T) {
	today := time.Now().Format("2006-01-02")
	if GetYMDFormat(0) != today {
		t.Errorf("GetYMDFormat(0) == %v, expected %v", GetYMDFormat(0), today)
	}
}

func TestTimeStampToTime(t *testing.T) {
	ts := time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local).Unix()
	res := TimeStampToTime(ts)
	// 只校验日期部分，避免时区差异导致的小时偏移
	if res[:10] != "2023-01-01" {
		t.Errorf("TimeStampToTime error: %v", res)
	}
}

func TestGetTimeRange(t *testing.T) {
	start, end := GetTimeRange(0)
	if len(start) != 19 || len(end) != 19 {
		t.Errorf("GetTimeRange error: %v, %v", start, end)
	}
}

func TestGetTodayRemainSeconds(t *testing.T) {
	remain := GetTodayRemainSeconds()
	if remain <= 0 || remain > 86400 {
		t.Errorf("GetTodayRemainSeconds error: %v", remain)
	}
}

func TestTruncateTimeToStartAndEnd(t *testing.T) {
	now := time.Now()
	start := TruncateTimeToStart(now)
	end := TruncateTimeToEnd(now)
	if end-start != 86399 {
		t.Errorf("TruncateTimeToStart/End error: %v, %v", start, end)
	}
}
