package httprequest

import (
	"context"
	"io"
	"testing"
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
