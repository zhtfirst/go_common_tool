package tool

import (
	"fmt"
	"testing"
)

func TestSeliceReverse(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"1", "2", "3", "4"}, []string{"4", "3", "2", "1"}},
	}

	for _, c := range cases {
		result := SeliceReverse(c.input)
		for i, v := range result {
			if v != c.expected[i] {
				t.Errorf("SeliceReverse(%v) == %v, expected %v", c.input, result, c.expected)
			}
		}
	}
}

func TestSeliceUnique(t *testing.T) {
	cases := []struct {
		input    []string
		expected []string
	}{
		{[]string{"a", "b", "a", "c"}, []string{"a", "b", "c"}},
		{[]string{"1", "2", "2", "3", "1"}, []string{"1", "2", "3"}},
	}

	for _, c := range cases {
		result := SeliceUnique(c.input)
		if len(result) != len(c.expected) {
			t.Errorf("SeliceUnique(%v) == %v, expected %v", c.input, result, c.expected)
		}
		for _, v := range result {
			found := false
			for _, ev := range c.expected {
				if v == ev {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("SeliceUnique(%v) == %v, expected %v", c.input, result, c.expected)
			}
		}
	}
}

// 测试 RandArray 洗牌算法
func TestRandArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	RandArray(arr)
	fmt.Println(arr)
	if len(arr) != 10 {
		t.Errorf("RandArray length changed, got %v", arr)
	}
	// 检查元素是否都在
	m := map[int]bool{}
	for _, v := range arr {
		m[v] = true
	}
	for i := 1; i <= 5; i++ {
		if !m[i] {
			t.Errorf("RandArray missing element %d", i)
		}
	}
}

// 测试 GetRandArrayValue
func TestGetRandArrayValue(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}
	val := GetRandArrayValue(arr)
	found := false
	for _, v := range arr {
		if v == val {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("GetRandArrayValue returned value not in array: %v", val)
	}
}

// 测试 GetPercentValue
func TestGetPercentValue(t *testing.T) {
	pct := []int{1, 2, 7}
	// 统计出现次数
	count := make([]int, len(pct))
	total := 10000
	for i := 0; i < total; i++ {
		idx := GetPercentValue(pct)
		if idx < 0 || idx >= len(pct) {
			t.Errorf("GetPercentValue returned out of range: %v", idx)
		}
		count[idx]++
	}
	// 检查比例大致正确
	if count[2] < count[1] || count[1] < count[0] {
		t.Errorf("GetPercentValue distribution error: %v", count)
	}
}

// 测试 RandomElement
func TestRandomElement(t *testing.T) {
	arr := []int{100, 200, 300}
	val, err := RandomElement(arr)
	if err != nil {
		t.Errorf("RandomElement error: %v", err)
	}
	found := false
	for _, v := range arr {
		if v == val {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("RandomElement returned value not in array: %v", val)
	}

	// 空切片测试
	_, err = RandomElement([]int{})
	if err == nil {
		t.Error("RandomElement should return error for empty slice")
	}
}
