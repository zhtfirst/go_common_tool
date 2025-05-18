package tool

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
}

// SeliceReverse 逆序输出字符串切片
// 参数：array 待逆序的字符串切片
// 返回值：逆序后的字符串切片（原地修改）
func SeliceReverse(array []string) []string {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-1-i] = array[len(array)-1-i], array[i]
	}
	return array
}

// SeliceUnique 字符串切片去重
// 参数：array 待去重的字符串切片
// 返回值：去重后的字符串切片（顺序不保证）
func SeliceUnique(array []string) []string {
	m := make(map[string]bool)
	for _, v := range array {
		m[v] = true
	}
	var result []string
	for k := range m {
		result = append(result, k)
	}
	return result
}

// RandArray 对整型切片进行原地随机乱序（洗牌算法）
// 参数：array 需要乱序的整型切片
func RandArray(array []int) {
	num := len(array)
	for i := num - 1; i > 0; i-- {
		temp := GetRandNum(0, i)
		array[i], array[temp] = array[temp], array[i]
	}
}

// GetRandArrayValue 从整型切片中随机返回一个元素
// 参数：array 整型切片
// 返回值：随机选中的元素
func GetRandArrayValue(array []int) int {
	num := len(array)
	idx := GetRandNum(0, num-1)
	return array[idx]
}

// GetPercentValue 按权重数组随机返回下标
// 参数：pct 权重数组（每个元素为权重）
// 返回值：根据权重随机选中的下标
// 例如：pct=[2,3,5]，返回0概率20%，1概率30%，2概率50%
func GetPercentValue(pct []int) int {
	percent := []int{}
	for idx, num := range pct {
		for m := 0; m < num; m++ {
			percent = append(percent, idx)
		}
	}

	RandArray(percent)
	return GetRandArrayValue(percent)
}

// RandomElement 从整型切片中随机取出一个元素
// 参数：slice 整型切片
// 返回值：(1) 随机选中的元素 (2) 错误信息（切片为空时返回错误）
func RandomElement(slice []int) (int, error) {
	length := len(slice)
	if length == 0 {
		return 0, fmt.Errorf("slice is empty")
	}

	index := rand.Intn(length) // 生成0到length-1之间的随机数
	return slice[index], nil
}
