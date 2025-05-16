package tool

import "testing"

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
