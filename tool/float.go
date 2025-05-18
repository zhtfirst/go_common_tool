/*
 * @Author: gavin zhtfirst@163.com
 * @Date: 2023-12-11 17:58:56
 * @LastEditors: gavin zhtfirst@163.com
 * @LastEditTime: 2025-05-18 10:10:29
 * @FilePath: /go_common_tool/tool/float.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tool

import (
	"fmt"
	"math"
	"strconv"
)

// RoundFloat 将浮点数四舍五入到指定小数位
//
// 参数:
//   - value: 需要处理的浮点数
//   - places: 要保留的小数位数
//
// 返回:
//   - 保留指定小数位后的浮点数
func RoundFloat(value float64, places int) float64 {
	if places < 0 {
		places = 0
	}

	// 使用数学方法实现四舍五入，性能更好
	factor := math.Pow10(places)

	return math.Round(value*factor+math.Copysign(1e-9, value)) / factor
}

// FormatFloat 使用字符串格式化方式保留指定位数的小数
//
// 参数:
//   - value: 需要处理的浮点数
//   - places: 要保留的小数位数
//
// 返回:
//   - 保留指定小数位的浮点数
//   - 可能的错误信息
func FormatFloat(value float64, places int) (float64, error) {
	if places < 0 {
		places = 0
	}

	// 修正浮点误差
	value = value + math.Copysign(1e-9, value)

	format := fmt.Sprintf("%%.%df", places)
	s := fmt.Sprintf(format, value)

	return strconv.ParseFloat(s, 64)
}
