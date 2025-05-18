package tool

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpPostExample(t *testing.T) {
	// 启动一个本地 http server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("期望POST方法，实际为: %s", r.Method)
		}
		var data map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			t.Errorf("POST请求体解析失败: %v", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	}))
	defer ts.Close()

	// 构造参数
	apiURL := ts.URL
	endpoint := "test"
	data := map[string]interface{}{"foo": "bar"}
	headers := map[string]string{"X-Test": "1"}

	resp, err := HttpPostExample(context.Background(), apiURL, endpoint, data, headers)
	if err != nil {
		t.Fatalf("HttpPostExample 返回错误: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("HttpPostExample 状态码错误: %d", resp.StatusCode())
	}
	var respData map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &respData); err != nil {
		t.Errorf("响应体解析失败: %v", err)
	}
	if ok, exists := respData["ok"]; !exists || ok != true {
		t.Errorf("响应体内容错误: %v", respData)
	}
}

func TestHttpGetExample(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("期望GET方法，实际为: %s", r.Method)
		}
		if r.URL.Query().Get("foo") != "bar" {
			t.Errorf("URL参数错误: %v", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ok":true}`))
	}))
	defer ts.Close()

	apiURL := ts.URL
	endpoint := "test"
	queryParams := map[string]string{"foo": "bar"}
	headers := map[string]string{"X-Test": "1"}

	resp, err := HttpGetExample(context.Background(), apiURL, endpoint, queryParams, headers)
	if err != nil {
		t.Fatalf("HttpGetExample 返回错误: %v", err)
	}
	if resp.StatusCode() != http.StatusOK {
		t.Errorf("HttpGetExample 状态码错误: %d", resp.StatusCode())
	}
	var respData map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &respData); err != nil {
		t.Errorf("响应体解析失败: %v", err)
	}
	if ok, exists := respData["ok"]; !exists || ok != true {
		t.Errorf("响应体内容错误: %v", respData)
	}
}
