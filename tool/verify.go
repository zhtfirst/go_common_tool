package tool

import "regexp"

// 校验工具包

// VerifyEmailFormat 校验邮箱格式是否合法
//
// 参数：
//   - email: 待校验的邮箱地址字符串
//
// 返回值：
//   - bool: 如果邮箱格式合法返回 true，否则返回 false
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
