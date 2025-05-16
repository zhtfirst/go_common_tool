package tool

import "testing"

func TestShuffleString(t *testing.T) {
	input := "abcdef"
	result := ShuffleString(input)
	if len(result) != len(input) {
		t.Errorf("ShuffleString(%v) == %v, expected length %v", input, result, len(input))
	}
}

func TestMd5(t *testing.T) {
	input := "hello"
	expected := "5d41402abc4b2a76b9719d911017c592"
	result := Md5(input)
	if result != expected {
		t.Errorf("Md5(%v) == %v, expected %v", input, result, expected)
	}
}

func TestStringArrayUnique(t *testing.T) {
	input := []string{"a", "b", "a", "c"}
	expected := []string{"a", "b", "c"}
	result := StringArrayUnique(input)
	if len(result) != len(expected) {
		t.Errorf("StringArrayUnique(%v) == %v, expected %v", input, result, expected)
	}
}

func TestCondExprInt(t *testing.T) {
	if CondExprInt(true, 1, 2) != 1 {
		t.Error("CondExprInt(true, 1, 2) != 1")
	}
	if CondExprInt(false, 1, 2) != 2 {
		t.Error("CondExprInt(false, 1, 2) != 2")
	}
}

func TestCondExprString(t *testing.T) {
	if CondExprString(true, "a", "b") != "a" {
		t.Error("CondExprString(true, \"a\", \"b\") != \"a\"")
	}
	if CondExprString(false, "a", "b") != "b" {
		t.Error("CondExprString(false, \"a\", \"b\") != \"b\"")
	}
}

func TestMaskStringNum(t *testing.T) {
	input := "1234567890"
	expected := "123****890"
	result := MaskStringNum(input, 4)
	if result != expected {
		t.Errorf("MaskStringNum(%v, 4) == %v, expected %v", input, result, expected)
	}
}

func TestTrimAndCombine(t *testing.T) {
	input := "  a,  b ,c , ,d  "
	expected := "a,b,c,d"
	result := TrimAndCombine(input)
	if result != expected {
		t.Errorf("TrimAndCombine(%v) == %v, expected %v", input, result, expected)
	}
}
