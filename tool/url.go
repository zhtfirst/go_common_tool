package tool

import (
	"github.com/spf13/cast"
	"net/url"
)

// MapToQuery
//
//	@Description: map转url query
//	@param parms
//	@return string
func MapToQuery(parms map[string]any) string {
	values := url.Values{}

	for k, v := range parms {
		values.Add(k, cast.ToString(v))
	}

	return values.Encode()
}

// HttpParams2String 拼接url链接
func HttpParams2String(url string, params map[string]string) string {
	times := len(params)
	i := 1
	res := url + "?"
	for k, v := range params {
		if i != times {
			res = res + k + "=" + v + "&"
			i = i + 1
		} else {
			res = res + k + "=" + v
		}
	}
	return res
}
