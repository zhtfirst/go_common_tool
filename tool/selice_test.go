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

// 测试 RandomElement
func TestRandomElement(t *testing.T) {
	arr := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	val, err := RandomElement(arr)
	fmt.Println(val)
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
