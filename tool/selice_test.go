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

func TestReverseSlice(t *testing.T) {
	// 测试整型切片
	t.Run("整型切片", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		ReverseSlice(slice)

		for i, v := range slice {
			if v != expected[i] {
				t.Errorf("ReverseSlice([]int{1, 2, 3, 4, 5}) = %v, 期望 %v", slice, expected)
				break
			}
		}
	})

	// 测试字符串切片
	t.Run("字符串切片", func(t *testing.T) {
		slice := []string{"a", "b", "c", "d"}
		expected := []string{"d", "c", "b", "a"}
		ReverseSlice(slice)

		for i, v := range slice {
			if v != expected[i] {
				t.Errorf("ReverseSlice([]string{\"a\", \"b\", \"c\", \"d\"}) = %v, 期望 %v", slice, expected)
				break
			}
		}
	})

	// 测试空切片
	t.Run("空切片", func(t *testing.T) {
		slice := []float64{}
		ReverseSlice(slice)
		if len(slice) != 0 {
			t.Errorf("ReverseSlice(空切片)后长度应为0，实际为 %d", len(slice))
		}
	})

	// 测试单元素切片
	t.Run("单元素切片", func(t *testing.T) {
		slice := []int{42}
		expected := []int{42}
		ReverseSlice(slice)

		if slice[0] != expected[0] {
			t.Errorf("ReverseSlice([]int{42}) = %v, 期望 %v", slice, expected)
		}
	})

	// 测试自定义类型切片
	t.Run("结构体切片", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		slice := []Person{
			{"张三", 25},
			{"李四", 30},
			{"王五", 35},
		}

		expected := []Person{
			{"王五", 35},
			{"李四", 30},
			{"张三", 25},
		}

		ReverseSlice(slice)

		for i, v := range slice {
			if v.Name != expected[i].Name || v.Age != expected[i].Age {
				t.Errorf("ReverseSlice失败：%v", slice)
				break
			}
		}
	})
}
