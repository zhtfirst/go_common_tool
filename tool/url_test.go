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
