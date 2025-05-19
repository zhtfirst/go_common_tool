/*
 * @Author: gavin
 * @Date: 2025-04-24 17:00:30
 * @LastEditors: gavin
 * @LastEditTime: 2025-05-19 16:09:38
 * @FilePath: /httprequest/httprequest_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package httprequest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestHttp
//
//	@Description: 测试http请求
//	@param t
func TestHttp(t *testing.T) {
	req := struct {
		Node   string `form:"node"`
		ID     int    `form:"id"`
		Header string `header:"X-Header"`
	}{
		Node:   "foo",
		ID:     10,
		Header: "my-header",
	}

	response, err := Http(context.Background(), "http://www.baidu.com", "GET", req, WithTimeout(12), WithServiceName("test"))

	if err != nil {
		t.Error(err)
		return
	}

	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))
}

// TestHttpPost
//
//	@Description: 测试POST请求
//	@param t
func TestHttpPost(t *testing.T) {
	// 创建测试服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// 读取请求体
		body, _ := io.ReadAll(r.Body)
		var data map[string]interface{}
		json.Unmarshal(body, &data)

		// 返回响应
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":   "success",
			"method":   "POST",
			"received": data,
		})
	}))
	defer server.Close()

	// 创建请求数据
	req := struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	}{
		Name:  "test-post",
		Value: 42,
	}

	response, err := Http(context.Background(), server.URL, "POST", req, WithTimeout(5))
	if err != nil {
		t.Fatal(err)
	}

	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	if result["method"] != "POST" {
		t.Errorf("Expected method POST, got %v", result["method"])
	}
}

// TestHttpPut
//
//	@Description: 测试PUT请求
//	@param t
func TestHttpPut(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT request, got %s", r.Method)
		}

		// 检查请求头
		headerValue := r.Header.Get("X-Custom-Header")
		if headerValue != "custom-value" {
			t.Errorf("Expected X-Custom-Header to be custom-value, got %s", headerValue)
		}

		// 读取请求体
		body, _ := io.ReadAll(r.Body)
		var data map[string]interface{}
		json.Unmarshal(body, &data)

		// 返回响应
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":   "updated",
			"method":   "PUT",
			"received": data,
		})
	}))
	defer server.Close()

	// 创建请求数据
	req := struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Header string `header:"X-Custom-Header"`
	}{
		ID:     123,
		Name:   "updated-item",
		Header: "custom-value",
	}

	response, err := Http(context.Background(), server.URL, "PUT", req)
	if err != nil {
		t.Fatal(err)
	}

	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	if result["method"] != "PUT" {
		t.Errorf("Expected method PUT, got %v", result["method"])
	}
}

// TestHttpDelete
//
//	@Description: 测试DELETE请求
//	@param t
func TestHttpDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("Expected DELETE request, got %s", r.Method)
		}

		// 检查URL参数
		id := r.URL.Query().Get("id")
		if id != "42" {
			t.Errorf("Expected id=42, got %s", id)
		}

		// 返回响应
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "deleted",
			"method": "DELETE",
			"id":     id,
		})
	}))
	defer server.Close()

	// 创建请求数据
	req := struct {
		ID int `form:"id"`
	}{
		ID: 42,
	}

	response, err := Http(context.Background(), server.URL, "DELETE", req, WithServiceName("delete-test"))
	if err != nil {
		t.Fatal(err)
	}

	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	if result["method"] != "DELETE" {
		t.Errorf("Expected method DELETE, got %v", result["method"])
	}
}

// TestHttpPatch
//
//	@Description: 测试PATCH请求
//	@param t
func TestHttpPatch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("Expected PATCH request, got %s", r.Method)
		}

		// 读取请求体
		body, _ := io.ReadAll(r.Body)
		var data map[string]interface{}
		json.Unmarshal(body, &data)

		// 返回响应
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":   "patched",
			"method":   "PATCH",
			"received": data,
		})
	}))
	defer server.Close()

	// 创建请求数据
	req := struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}{
		Field: "status",
		Value: "active",
	}

	response, err := Http(context.Background(), server.URL, "PATCH", req, WithTimeout(3))
	if err != nil {
		t.Fatal(err)
	}

	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	if result["method"] != "PATCH" {
		t.Errorf("Expected method PATCH, got %v", result["method"])
	}
}

// TestHttpWithMixedParams
//
//	@Description: 测试混合参数类型的请求
//	@param t
func TestHttpWithMixedParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查请求头
		authHeader := r.Header.Get("Authorization")
		traceHeader := r.Header.Get("X-Trace-ID")

		// 检查URL参数
		queryID := r.URL.Query().Get("id")
		queryPage := r.URL.Query().Get("page")

		// 读取请求体
		body, _ := io.ReadAll(r.Body)
		var data map[string]interface{}
		if len(body) > 0 {
			json.Unmarshal(body, &data)
		}

		// 返回响应
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"method": r.Method,
			"headers": map[string]string{
				"Authorization": authHeader,
				"X-Trace-ID":    traceHeader,
			},
			"query": map[string]string{
				"id":   queryID,
				"page": queryPage,
			},
			"body": data,
		})
	}))
	defer server.Close()

	// 创建请求数据
	req := struct {
		ID        int    `form:"id"`
		Page      int    `form:"page"`
		Auth      string `header:"Authorization"`
		TraceID   string `header:"X-Trace-ID"`
		Data      string `json:"data"`
		Timestamp int64  `json:"timestamp"`
	}{
		ID:        100,
		Page:      1,
		Auth:      "Bearer token123",
		TraceID:   "trace-abc-123",
		Data:      "test-data",
		Timestamp: 1625097600,
	}

	response, err := Http(context.Background(), server.URL, "POST", req)
	if err != nil {
		t.Fatal(err)
	}

	body, _ := io.ReadAll(response.Body)
	t.Log(string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	// 验证结果
	headers, ok := result["headers"].(map[string]interface{})
	if !ok {
		t.Fatal("Headers not found in response")
	}

	if headers["Authorization"] != "Bearer token123" {
		t.Errorf("Expected Authorization header Bearer token123, got %v", headers["Authorization"])
	}

	if headers["X-Trace-ID"] != "trace-abc-123" {
		t.Errorf("Expected X-Trace-ID header trace-abc-123, got %v", headers["X-Trace-ID"])
	}
}

// TestHttpWithInvalidMethod
//
//	@Description: 测试无效的HTTP方法
//	@param t
func TestHttpWithInvalidMethod(t *testing.T) {
	req := struct {
		Name string `json:"name"`
	}{
		Name: "test",
	}

	_, err := Http(context.Background(), "http://example.com", "INVALID_METHOD", req)

	if err == nil {
		t.Error("Expected error for invalid HTTP method, got nil")
	}

	if err != noSupportMethod {
		t.Errorf("Expected noSupportMethod error, got %v", err)
	}
}

// TestHttpHead
//
//	@Description: 测试HEAD请求
//	@param t
func TestHttpHead(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodHead {
			t.Errorf("Expected HEAD request, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Test-Header", "test-value")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	response, err := Http(context.Background(), server.URL, "HEAD", nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	contentType := response.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	testHeader := response.Header.Get("X-Test-Header")
	if testHeader != "test-value" {
		t.Errorf("Expected X-Test-Header test-value, got %s", testHeader)
	}

	// HEAD请求不应该有响应体
	body, _ := io.ReadAll(response.Body)
	if len(body) > 0 {
		t.Errorf("Expected empty body for HEAD request, got %d bytes", len(body))
	}
}

// TestHttpWithTimeout
//
//	@Description: 测试带超时的HTTP请求
//	@param t
func TestHttpWithTimeout(t *testing.T) {
	// 创建一个会延迟响应的服务器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 延迟3秒响应
		time.Sleep(3 * time.Second)
		fmt.Fprintln(w, "Response after delay")
	}))
	defer server.Close()

	// 使用1秒的超时设置
	_, err := Http(context.Background(), server.URL, "GET", nil, WithTimeout(1))

	// 应该返回超时错误
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}

// TestHttpWithContextCancel
//
//	@Description: 测试带取消的HTTP请求
//	@param t
func TestHttpWithContextCancel(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 延迟2秒响应
		time.Sleep(2 * time.Second)
		fmt.Fprintln(w, "Response after delay")
	}))
	defer server.Close()

	// 创建可取消的上下文
	ctx, cancel := context.WithCancel(context.Background())

	// 在500毫秒后取消请求
	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()

	_, err := Http(ctx, server.URL, "GET", nil)

	// 应该返回上下文取消错误
	if err == nil {
		t.Error("Expected context canceled error, got nil")
	}
}

// TestHttpWithRetry
//
//	@Description: 测试HTTP请求重试
//	@param t
func TestHttpWithRetry(t *testing.T) {
	attempts := 0
	maxAttempts := 3

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < maxAttempts {
			// 前两次请求返回503错误
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		// 第三次请求成功
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Success after retry")
	}))
	defer server.Close()

	// 自定义重试逻辑
	var response *http.Response
	var err error

	for i := 0; i < maxAttempts; i++ {
		response, err = Http(context.Background(), server.URL, "GET", nil)
		if err != nil {
			continue
		}

		if response.StatusCode == http.StatusOK {
			break
		}

		// 关闭响应体以避免资源泄漏
		if response != nil && response.Body != nil {
			response.Body.Close()
		}

		// 等待一段时间后重试
		time.Sleep(100 * time.Millisecond)
	}

	if err != nil {
		t.Fatalf("Failed after %d attempts: %v", maxAttempts, err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	body, _ := io.ReadAll(response.Body)
	if string(body) != "Success after retry\n" {
		t.Errorf("Unexpected response body: %s", string(body))
	}
}

// TestHttpWithDifferentContentTypes
//
//	@Description: 测试不同Content-Type的HTTP请求
//	@param t
func TestHttpWithDifferentContentTypes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"received_content_type": "%s"}`, contentType)
	}))
	defer server.Close()

	// 1. 测试JSON请求
	jsonReq := struct {
		Data string `json:"data"`
	}{
		Data: "json-data",
	}

	jsonResp, err := Http(context.Background(), server.URL, "POST", jsonReq)
	if err != nil {
		t.Fatal(err)
	}

	jsonBody, _ := io.ReadAll(jsonResp.Body)
	t.Logf("JSON request response: %s", string(jsonBody))

	// 2. 测试表单请求
	formReq := struct {
		Data string `form:"data"`
	}{
		Data: "form-data",
	}

	formResp, err := Http(context.Background(), server.URL, "POST", formReq)
	if err != nil {
		t.Fatal(err)
	}

	formBody, _ := io.ReadAll(formResp.Body)
	t.Logf("Form request response: %s", string(formBody))
}

// TestHttpWithLargeResponse
//
//	@Description: 测试大响应体的HTTP请求
//	@param t
func TestHttpWithLargeResponse(t *testing.T) {
	// 生成大约1MB的响应数据
	largeData := make([]byte, 1024*1024)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(largeData)
	}))
	defer server.Close()

	response, err := Http(context.Background(), server.URL, "GET", nil)
	if err != nil {
		t.Fatal(err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	if len(body) != len(largeData) {
		t.Errorf("Expected response body length %d, got %d", len(largeData), len(body))
	}

	// 检查前100个字节是否匹配
	for i := 0; i < 100; i++ {
		if body[i] != largeData[i] {
			t.Errorf("Response body mismatch at position %d", i)
			break
		}
	}
}
