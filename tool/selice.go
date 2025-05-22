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

// RandomElement 从整型切片中随机取出一个元素
// 参数：slice 整型切片
// 返回值：(1) 随机选中的元素 (2) 错误信息（切片为空时返回错误）
func RandomElement(slice []int) (int, error) {
	length := len(slice)
	if length == 0 {
		return 0, fmt.Errorf("slice is empty")
	}

	index := GetRandNum(0, length-1) // 生成0到length-1之间的随机数
	return slice[index], nil
}

// ReverseSlice 反转任意类型的切片
// 函数没有返回值是因为它采用了原地修改的设计模式。在 Go 中，切片是引用类型，当你将切片传递给函数时，函数内对切片的修改会直接反映到原始切片上。
func ReverseSlice[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
