/*
 * @Author: gavin zhtfirst@163.com
 * @Date: 2025-05-18 10:40:40
 * @LastEditors: gavin zhtfirst@163.com
 * @LastEditTime: 2025-05-18 10:56:20
 * @FilePath: /go_common_tool/tool/http.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tool

import (
	"context"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

/*
* go-resty v2 文档 https://github.com/go-resty/resty/blob/v2/README.md
* go-resty v3 文档 https://resty.dev/
 */

// HttpPostExample 展示如何使用resty发送POST请求
// ctx: 上下文，用于日志记录或请求取消
// apiURL: API基础URL
// endpoint: 具体接口路径
// data: 请求体数据
// headers: 可选的额外请求头
func HttpPostExample(ctx context.Context, apiURL, endpoint string, data interface{}, headers map[string]string) (*resty.Response, error) {
	// 构建完整URL
	url := fmt.Sprintf("%s/%s", apiURL, endpoint)

	// 创建resty客户端
	client := resty.New()

	// 设置基本配置
	client.SetTimeout(5 * time.Second) // 设置超时时间
	client.SetRetryCount(3)            // 设置重试次数

	// 准备请求
	request := client.R().
		SetContext(ctx).
		SetHeader("Content-Type", "application/json").
		SetBody(data)

	// 添加自定义头部
	for k, v := range headers {
		request.SetHeader(k, v)
	}

	// 发送POST请求
	resp, err := request.Post(url)

	// 记录请求日志
	fmt.Printf("POST请求: %s, 数据: %v\n", url, data)
	fmt.Printf("响应状态: %d, 响应体: %s, 错误: %v\n", resp.StatusCode(), resp.String(), err)

	return resp, err
}

// HttpGetExample 展示如何使用resty发送GET请求
// ctx: 上下文，用于日志记录或请求取消
// apiURL: API基础URL
// endpoint: 具体接口路径
// queryParams: URL查询参数
// headers: 可选的额外请求头
func HttpGetExample(ctx context.Context, apiURL, endpoint string, queryParams map[string]string, headers map[string]string) (*resty.Response, error) {
	// 构建完整URL
	url := fmt.Sprintf("%s/%s", apiURL, endpoint)

	// 创建resty客户端
	client := resty.New()

	// 设置基本配置
	client.SetTimeout(5 * time.Second) // 设置超时时间
	client.SetRetryCount(2)            // 设置重试次数

	// 准备请求
	request := client.R().
		SetContext(ctx).
		SetQueryParams(queryParams)

	// 添加自定义头部
	for k, v := range headers {
		request.SetHeader(k, v)
	}

	// 发送GET请求
	resp, err := request.Get(url)

	// 记录请求日志
	fmt.Printf("GET请求: %s, 参数: %v\n", url, queryParams)
	fmt.Printf("响应状态: %d, 响应体: %s, 错误: %v\n", resp.StatusCode(), resp.String(), err)

	return resp, err
}
