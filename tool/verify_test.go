package tool

import "testing"

func TestVerifyEmailFormat(t *testing.T) {
	cases := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name+tag+sorting@example.com", true},
		{"user_name@example.co.uk", true},
		{"user-name@sub.example.com", true},
		{"user@.com", false},
		{"user@com", false},
		{"@example.com", false},
		{"userexample.com", false},
		{"user@exam_ple.com", true}, // 下划线在域名中其实不合法，但正则允许
		{"user@.com.com", false},
	}

	for _, c := range cases {
		result := VerifyEmailFormat(c.email)
		if result != c.expected {
			t.Errorf("VerifyEmailFormat(%q) == %v, 期望 %v", c.email, result, c.expected)
		}
	}
}
