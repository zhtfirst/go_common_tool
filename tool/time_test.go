/*
 * @Author: gavin v_zhangtao15@tal.com
 * @Date: 2025-05-16 18:04:30
 * @LastEditors: gavin v_zhangtao15@tal.com
 * @LastEditTime: 2025-05-16 18:09:24
 * @FilePath: /go_common_tool/tool/time_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tool

import (
	"context"
	"testing"
	"time"
)

func TestDayDiff(t *testing.T) {
	days := DayDiff("2023-01-01", "2023-01-10")
	if days != 9 {
		t.Errorf("DayDiff(\"2023-01-01\", \"2023-01-10\") == %v, expected 9", days)
	}
}

func TestDateFormat(t *testing.T) {
	formatted := DateFormat("2023-01-01T15:04:05Z")
	if formatted != "2023-01-01" {
		t.Errorf("DateFormat(\"2023-01-01T15:04:05Z\") == %v, expected \"2023-01-01\"", formatted)
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
	expected := time.Now().Format("2006-01-02 15:04:05")
	if now != expected {
		t.Errorf("GetNowTime() == %v, expected %v", now, expected)
	}
}

func TestFormatTimeToString(t *testing.T) {
	t1 := time.Date(2023, 1, 1, 15, 4, 0, 0, time.UTC)
	formatted, err := FormatTimeToString(t1)
	if err != nil || formatted != "2023-01-01 15:04" {
		t.Errorf("FormatTimeToString(t1) == %v, %v, expected \"2023-01-01 15:04\", nil", formatted, err)
	}
}

func TestRemoveTimezone(t *testing.T) {
	ctx := context.TODO()
	timeStr := "2023-01-01 15:04:05 +0000 UTC"
	result := RemoveTimezone(ctx, timeStr)
	expected := "2023-01-01 15:04:05"
	if result != expected {
		t.Errorf("RemoveTimezone(ctx, %v) == %v, expected %v", timeStr, result, expected)
	}
}
