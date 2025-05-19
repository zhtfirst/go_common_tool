package tool

import (
	"testing"
)

func TestGetRandNum_NormalRange(t *testing.T) {
	min, max := 1, 10
	for i := 0; i < 10; i++ {
		n := GetRandNum(min, max)
		if n < min || n >= max {
			t.Errorf("GetRandNum(%d, %d) = %d, 超出范围", min, max, n)
		}
	}
}

func TestGetRandNum_MinEqualsMax(t *testing.T) {
	min, max := 5, 5
	n := GetRandNum(min, max)
	if n != min {
		t.Errorf("GetRandNum(%d, %d) = %d, 期望值: %d", min, max, n, min)
	}
}

func TestGetRandNum_MinGreaterThanMax(t *testing.T) {
	min, max := 10, 1
	for i := 0; i < 100; i++ {
		n := GetRandNum(min, max)
		if n < max || n >= min {
			t.Errorf("GetRandNum(%d, %d) = %d, 超出范围", min, max, n)
		}
	}
}

func TestGetRandNum_NegativeRange(t *testing.T) {
	min, max := -10, -1
	for i := 0; i < 100; i++ {
		n := GetRandNum(min, max)
		if n < min || n >= max {
			t.Errorf("GetRandNum(%d, %d) = %d, 超出范围", min, max, n)
		}
	}
}

// TestGenerateSecureRandomString
//
//	@Description: 测试GenerateSecureRandomString
//	@param t
func TestGenerateSecureRandomString(t *testing.T) {
	type args struct {
		length int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				length: 10,
			},
		},
		{
			name: "test2",
			args: args{
				length: 20,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(GenerateSecureRandomString(tt.args.length))
		})
	}
}
