package jwt

// Token
// @Description: jwt token
type Token struct {
	AccessExpire int64          `json:"access_expire"`
	AccessSecret string         `json:"access_secret"`
	Payload      map[string]any `json:"payload,omitempty"`
}

// TokenResponse
// @Description: jwt token返回
type TokenResponse struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
	Iat   int64  `json:"iat"`
}

// TokenParseResponse
// @Description: jwt token解析返回
type TokenParseResponse struct {
	Valid   bool           `json:"valid"`   // 是否通过token解析验证
	Payload map[string]any `json:"payload"` // token解析后的数据payload
}
