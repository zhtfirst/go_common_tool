package tool

import "testing"

func TestHttpParams2String(t *testing.T) {
	url := "http://example.com"
	params := map[string]string{"a": "1", "b": "2"}
	expected := "http://example.com?a=1&b=2"
	result := HttpParams2String(url, params)
	if result != expected {
		t.Errorf("HttpParams2String(%v, %v) == %v, expected %v", url, params, result, expected)
	}
}

// TestMapToQuery
//
//	@Description: map转url query
//	@param t
func TestMapToQuery(t *testing.T) {
	// 正常情况测试
	parms := make(map[string]any)
	parms["key1"] = "value1"
	parms["key2"] = "value2"
	expected := "key1=value1&key2=value2"
	result := MapToQuery(parms)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// 空参数测试
	parms = make(map[string]any)
	expected = ""
	result = MapToQuery(parms)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// 非字符串值测试
	parms = make(map[string]any)
	parms["key1"] = 123
	expected = "key1=123"
	result = MapToQuery(parms)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
