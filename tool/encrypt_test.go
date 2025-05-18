/*
 * @Author: gavin zhtfirst@163.com
 * @Date: 2025-05-18 10:12:29
 * @LastEditors: gavin zhtfirst@163.com
 * @LastEditTime: 2025-05-18 10:15:39
 * @FilePath: /go_common_tool/tool/encrypt_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tool

import (
	"testing"
)

func TestSHA1(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"hello", "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"},
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
	}
	for _, c := range cases {
		got := SHA1(c.input)
		if got != c.expected {
			t.Errorf("SHA1(%q) = %q, want %q", c.input, got, c.expected)
		}
	}
}

func TestMd5String(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
	}
	for _, c := range cases {
		got := Md5String(c.input)
		if got != c.expected {
			t.Errorf("Md5String(%q) = %q, want %q", c.input, got, c.expected)
		}
	}
}

func TestMD5(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
	}
	for _, c := range cases {
		got := MD5(c.input)
		if got != c.expected {
			t.Errorf("MD5(%q) = %q, want %q", c.input, got, c.expected)
		}
	}
}

func TestSha512(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"hello", "9b71d224bd62f3785d96d46ad3ea3d73319bfbc2890caadae2dff72519673ca72323c3d99ba5c11d7c7acc6e14b8c5da0c4663475c2e5c3adef46f73bcdec043"},
		{"", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
	}
	for _, c := range cases {
		got := Sha512(c.input)
		if got != c.expected {
			t.Errorf("Sha512(%q) = %q, want %q", c.input, got, c.expected)
		}
	}
}

func TestComputeHmacSha256(t *testing.T) {
	cases := []struct {
		message  string
		secret   string
		expected string
	}{
		{"hello", "key", "9307B3B915EFB5171FF14D8CB55FBCC798C6C0EF1456D66DED1A6AA723A58B7B"},
	}
	for _, c := range cases {
		got := ComputeHmacSha256(c.message, c.secret)
		if got != c.expected {
			t.Errorf("ComputeHmacSha256(%q, %q) = %q, want %q", c.message, c.secret, got, c.expected)
		}
	}
}
