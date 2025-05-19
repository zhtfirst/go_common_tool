package httprequest

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"

	"github.com/zeromicro/go-zero/rest/httpc"
)

type Option func(*option)

// option
// @Description: http请求选项
type option struct {
	serviceName string
	timeout     int
}

// Http
//
//	@Description: http请求
//	@param ctx
//	@param url
//	@param method
//	@param req
//	@param opts
//	@return *http.Response
//	@return error
func Http(ctx context.Context, url, method string, req interface{}, opts ...Option) (*http.Response, error) {
	// 检查请求方法是否支持
	method, ok := checkRequestMethod(method)
	if !ok {
		return nil, noSupportMethod
	}

	// 构建请求
	opt := new(option)
	for _, o := range opts {
		o(opt)
	}

	timeout := getTimeout(opt)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()

	startTime := timex.Now()
	response, err := newService(getServiceName(opt)).Do(ctx, method, url, req)
	duration := timex.Since(startTime)
	if err != nil {
		return nil, err
	}

	logLevel := "httprequest." + method
	var body []byte
	if err != nil {
		logLevel += logLevel + ".error"
	} else {
		body, err = io.ReadAll(response.Body)
		if err != nil {
			if !strings.HasSuffix(logLevel, "error") {
				logLevel += logLevel + ".error"
			}
		}
	}

	requestParam, _ := json.Marshal(req)
	logx.WithContext(ctx).WithDuration(duration).Infow("",
		logx.Field("x_name", logLevel),
		logx.Field("x_action", url),
		logx.Field("x_param", string(requestParam)),
		logx.Field("x_response", Substring(string(body), responseLength)),
	)

	if len(body) > 0 {
		response.Body = io.NopCloser(strings.NewReader(string(body)))
	}

	return response, nil
}

// checkRequestMethod
//
//	@Description: 检查请求方法是否支持
//	@param method
//	@return string
//	@return bool
func checkRequestMethod(method string) (string, bool) {
	method = strings.ToUpper(method)

	support, ok := supportMethod[method]
	return method, ok && support
}

// WithServiceName
//
//	@Description: 设置服务名称
//	@param name
//	@return Option
func WithServiceName(name string) Option {
	return func(o *option) {
		o.serviceName = name
	}
}

// WithTimeout
//
//	@Description: 设置超时时间
//	@param timeout
//	@return Option
func WithTimeout(timeout int) Option {
	return func(o *option) {
		o.timeout = timeout
	}
}

// getTimeout
//
//	@Description: 获取超时时间
//	@param opt
//	@return int
func getTimeout(opt *option) int {
	if opt.timeout > 0 {
		return opt.timeout
	}

	return defaultTimeout
}

// getServiceName
//
//	@Description: 获取服务名称
//	@param opt
//	@return string
func getServiceName(opt *option) string {
	if opt.serviceName != "" {
		return opt.serviceName
	}

	return defaultServiceName
}

// newService
//
//	@Description: 创建httpc服务
//	@param serviceName
//	@return httpc.Service
func newService(serviceName string) httpc.Service {
	return httpc.NewService(serviceName)
}

// Substring
//
//	@Description: 截取字符串
//	@param s
//	@param l
//	@return string
func Substring(s string, l int) string {
	str := []rune(s)
	res := ""
	resCount := 0
	if len(str) <= l {
		return s
	}

	for _, r := range str {
		if resCount+1 > l {
			break
		}
		resCount++
		res += string(r)
	}
	return res
}
