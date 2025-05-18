/*
 * @Author: gavin zhtfirst@163.com
 * @Date: 2025-05-18 09:48:04
 * @LastEditors: gavin zhtfirst@163.com
 * @LastEditTime: 2025-05-18 09:51:20
 * @FilePath: /go_common_tool/tool/ip_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tool

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockInterfaceAddrs 用于模拟 net.InterfaceAddrs 的返回值
func mockInterfaceAddrs(addrs []net.Addr, err error) func() ([]net.Addr, error) {
	return func() ([]net.Addr, error) {
		return addrs, err
	}
}

func TestGetIntranetIp(t *testing.T) {
	// 备份原有的 interfaceAddrsFunc
	origFunc := interfaceAddrsFunc
	defer func() { interfaceAddrsFunc = origFunc }()

	// 1. 正常返回 IPv4 地址
	ip := GetIntranetIp()
	if ip == "" {
		t.Error("期望获取到内网IP地址，但返回空字符串")
	}

	// 2. 模拟 interfaceAddrsFunc 返回错误
	interfaceAddrsFunc = mockInterfaceAddrs(nil, assert.AnError)
	ip = GetIntranetIp()
	if ip != "" {
		t.Error("当获取地址出错时，期望返回空字符串")
	}

	// 3. 模拟只有回环地址
	loopback := &net.IPNet{IP: net.ParseIP("127.0.0.1"), Mask: net.CIDRMask(8, 32)}
	interfaceAddrsFunc = mockInterfaceAddrs([]net.Addr{loopback}, nil)
	ip = GetIntranetIp()
	if ip != "" {
		t.Error("只有回环地址时，期望返回空字符串")
	}
}
