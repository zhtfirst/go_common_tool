/*
 * @Author: gavin zhtfirst@163.com
 * @Date: 2023-06-27 13:29:36
 * @LastEditors: gavin zhtfirst@163.com
 * @LastEditTime: 2025-05-18 09:31:56
 * @FilePath: /go_common_tool/tool/rand.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
// Package tool 提供常用的工具函数集合
package tool

import (
	"math/rand"
	"time"
)

// init 在包初始化时设置随机数种子，确保每次运行时生成的随机数不同。

// GetRandNum 返回[min_num, max_num)区间内的一个随机整数。
//
// 参数：
//
//	min_num: 随机数的最小值（包含）
//	max_num: 随机数的最大值（不包含）
//
// 返回值：
//
//	int: 区间[min_num, max_num)内的随机整数；若min_num==max_num，返回min_num。
//	     若min_num>max_num，则自动交换两者顺序。
func GetRandNum(min_num int, max_num int) int {
	// 创建独立的随机源
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	if min_num == max_num {
		return min_num
	}

	if min_num > max_num {
		min_num, max_num = max_num, min_num
	}

	return r.Intn(max_num-min_num) + min_num
}
