package tool

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// ShuffleString 随机打乱字符串
func ShuffleString(s string) string {
	// 将字符串转换为rune类型的切片
	r := []rune(s)
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// 打乱切片
	for i := len(r) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		r[i], r[j] = r[j], r[i]
	}
	// 将切片转换为字符串
	shuffledS := string(r)
	return shuffledS
}

// Md5 md5加密
func Md5(str string) string {
	data := []byte(str)
	md5New := md5.New()
	md5New.Write(data)
	// hex转字符串
	md5String := hex.EncodeToString(md5New.Sum(nil))
	return md5String
}

// StringArrayUnique 字符串切片去重实现
func StringArrayUnique(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for i := 0; i < len(arr); i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}

// CondExprInt 三元表达式，Int类型
func CondExprInt(expr bool, a, b int) int {
	if expr {
		return a
	}
	return b
}

// CondExprString 三元表达式，string类型
func CondExprString(expr bool, a, b string) string {
	if expr {
		return a
	}
	return b
}

// MaskStringNum 函数用于将字符串的中间num位替换为星号
// 场景：手机号、身份证号、银行卡号、邮箱地址等加密展示
func MaskStringNum(str string, num int) string {
	if len(str) < num {
		return str // 不足位数，不做处理
	}

	// 计算星号开始替换的位置
	start := (len(str) - num) / 2
	// 计算星号结束的位置
	end := start + num

	// 创建一个rune切片来处理可能的多字节字符
	runes := []rune(str)
	// 将中间的四个字符替换为星号
	for i := start; i < end; i++ {
		runes[i] = '*'
	}
	return string(runes)
}

// TrimAndCombine 去除字符串两边的空格、中英文逗号、字符串中的多个空格和多个逗号转换成一个
func TrimAndCombine(str string) string {
	// 去除两端的空格和逗号（中英文）
	str = strings.TrimSpace(str)
	str = strings.Trim(str, ",")
	str = strings.Trim(str, "，")

	// 将多个连续空格替换为单个空格
	spaceRegex := regexp.MustCompile(`\s+`)
	str = spaceRegex.ReplaceAllString(str, " ")

	// 将中文逗号和英文逗号统一为英文逗号
	str = strings.ReplaceAll(str, "，", ",")

	// 将多个连续逗号替换为单个逗号
	commaRegex := regexp.MustCompile(`,+`)
	str = commaRegex.ReplaceAllString(str, ",")

	// 处理逗号+空格或空格+逗号的情况
	mixedRegex := regexp.MustCompile(`[,\s]+`)
	str = mixedRegex.ReplaceAllString(str, ",")

	// 再次去除两端可能出现的逗号
	str = strings.Trim(str, ",")

	return str
}

// reverseString
//
//	@Description: 字符串反转
//	@param str
//	@return string
func reverseString(str string) string {
	var result strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}
	return result.String()
}
