/*
 * @Author: gavin zhtfirst@163.com
 * @Date: 2025-05-18 09:03:29
 * @LastEditors: gavin zhtfirst@163.com
 * @LastEditTime: 2025-05-18 09:48:10
 * @FilePath: /go_common_tool/tool/ip.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tool

import "net"

// ip.go 提供获取本机内网IP地址的工具函数

// interfaceAddrsFunc 用于获取所有网络接口的地址，默认指向 net.InterfaceAddrs，便于测试时注入 mock。
var interfaceAddrsFunc = net.InterfaceAddrs

// GetIntranetIp 获取本机的内网IPv4地址。
// 如果获取失败或没有找到合适的IP，返回空字符串。
func GetIntranetIp() string {
	// 获取所有网络接口的地址信息
	addrs, err := interfaceAddrsFunc()
	if err != nil {
		// 获取失败，返回空字符串
		return ""
	}

	// 遍历所有地址，查找非回环的IPv4地址
	for _, address := range addrs {
		// 类型断言，判断 address 是否为 *net.IPNet，并且不是回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// 判断是否为IPv4地址
			if ipnet.IP.To4() != nil {
				// 返回第一个找到的内网IPv4地址
				return ipnet.IP.String()
			}
		}
	}

	// 未找到合适的IP，返回空字符串
	return ""
}
