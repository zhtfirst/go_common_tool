/*
 * @Author: gavin zhtfirst@163.com
 * @Date: 2025-05-18 08:43:44
 * @LastEditors: gavin zhtfirst@163.com
 * @LastEditTime: 2025-05-18 10:12:31
 * @FilePath: /go_common_tool/tool/encrypt.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tool

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"strings"
)

// 加密工具包，提供常用的哈希和HMAC算法封装

// SHA1 对输入字符串进行SHA-1哈希，返回十六进制字符串
// 参数：
//
//	text: 需要加密的字符串
//
// 返回值：
//
//	string: SHA-1哈希后的十六进制字符串
func SHA1(text string) string {
	// 将字符串转换为字节数组
	data := []byte(text)

	// 创建SHA-1哈希对象
	hash := sha1.New()

	// 计算哈希值
	hash.Write(data)
	hashValue := hash.Sum(nil)

	// 将哈希值转换为十六进制字符串
	hashString := hex.EncodeToString(hashValue)

	return hashString
}

// Md5String 对输入字符串进行MD5哈希，返回小写十六进制字符串
// 参数：
//
//	data: 需要加密的字符串
//
// 返回值：
//
//	string: MD5哈希后的小写十六进制字符串
func Md5String(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// MD5 对输入字符串进行MD5哈希，返回小写十六进制字符串
// 参数：
//
//	text: 需要加密的字符串
//
// 返回值：
//
//	string: MD5哈希后的小写十六进制字符串
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

// Sha512 对输入字符串进行SHA-512哈希，返回小写十六进制字符串
// 参数：
//
//	text: 需要加密的字符串
//
// 返回值：
//
//	string: SHA-512哈希后的小写十六进制字符串
func Sha512(text string) string {
	sha := sha512.New()
	sha.Write([]byte(text))
	return hex.EncodeToString(sha.Sum(nil))
}

// ComputeHmacSha256 使用HMAC-SHA256算法对消息进行加密，返回大写十六进制字符串
// 参数：
//
//	message: 需要加密的消息内容
//	secret: HMAC密钥
//
// 返回值：
//
//	string: HMAC-SHA256加密后的大写十六进制字符串
func ComputeHmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	sha := hex.EncodeToString(h.Sum(nil))
	return strings.ToUpper(sha)
}
