package tool

import "testing"

func TestMapMerge(t *testing.T) {
	cases := []struct {
		array1   map[string]interface{}
		array2   []map[string]interface{}
		expected map[string]interface{}
	}{
		{
			map[string]interface{}{"a": 1, "b": 2},
			[]map[string]interface{}{
				{"b": 3, "c": 4},
			},
			map[string]interface{}{"a": 1, "b": 3, "c": 4},
		},
		{
			map[string]interface{}{"x": "foo"},
			[]map[string]interface{}{
				{"y": "bar"}, {"z": "baz"},
			},
			map[string]interface{}{"x": "foo", "y": "bar", "z": "baz"},
		},
	}

	for _, c := range cases {
		result := MapMerge(c.array1, c.array2...)
		for k, v := range c.expected {
			if result[k] != v {
				t.Errorf("MapMerge(%v, %v) == %v, expected %v", c.array1, c.array2, result, c.expected)
			}
		}
	}
}
