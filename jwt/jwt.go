package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// New
//
//	@Description: 创建jwt token
//	@param accessSecret
//	@param accessExpire
//	@return *Token
func New(accessSecret string, accessExpire int64) *Token {
	return &Token{
		AccessSecret: accessSecret,
		AccessExpire: accessExpire,
	}
}

// SetPayload
//
//	@Description: 设置jwt token payload
//	@receiver t
//	@param payload
//	@return *Token
func (t *Token) SetPayload(payload map[string]interface{}) *Token {
	t.Payload = payload

	return t
}

// GenerateToken
//
//	@Description: 生成jwt token
//	@receiver t
//	@return *TokenResponse
//	@return error
func (t *Token) GenerateToken() (*TokenResponse, error) {
	claims := t.getClaims()
	token, err := t.generateToken(claims)
	if err != nil {
		return nil, err
	}

	var (
		exp, _ = claims.GetExpirationTime()
		iat, _ = claims.GetIssuedAt()
	)

	return &TokenResponse{
		Token: token,
		Exp:   exp.Unix(),
		Iat:   iat.Unix(),
	}, nil
}

// Parse
//
//	@Description: 解析jwt token
//	@receiver t
//	@param tokenString
//	@return *TokenParseResponse
//	@return error
func (t *Token) Parse(tokenString string) (*TokenParseResponse, error) {
	jwtToken, err := t.parseJwtToken(tokenString)
	if err != nil {
		return nil, err
	}
	claims := jwtToken.Claims.(jwt.MapClaims)

	payload := claims["payload"]
	rv := &TokenParseResponse{
		Valid: true,
	}
	if payload != nil {
		rv.Payload = payload.(map[string]interface{})
	}

	return rv, nil
}

// parseJwtToken
//
//	@Description: 解析jwt token
//	@receiver t
//	@param tokenString
//	@return *jwt.Token
//	@return error
func (t *Token) parseJwtToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.AccessSecret), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
}

// getClaims
//
//	@Description: 获取jwt token claims
//	@receiver t
//	@return Claims
func (t *Token) getClaims() Claims {
	iat := jwt.NewNumericDate(time.Now())
	return Claims{
		IssuedAt:  iat,
		ExpiresAt: jwt.NewNumericDate(iat.Add(time.Duration(t.AccessExpire) * time.Second)),
		Payload:   t.Payload,
	}
}

// generateToken
//
//	@Description: 生成jwt token
//	@receiver t
//	@param claims
//	@return string
//	@return error
func (t *Token) generateToken(claims jwt.Claims) (string, error) {
	jwtToken := jwt.Token{
		Method: jwt.SigningMethodHS256,
		Header: map[string]interface{}{
			"typ": "JWT",
			"alg": jwt.SigningMethodHS256.Alg(),
		},
		Claims: claims,
		Valid:  true,
	}
	return jwtToken.SignedString([]byte(t.AccessSecret))
}
