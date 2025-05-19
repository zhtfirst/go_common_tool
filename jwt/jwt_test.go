package jwt

import "testing"

// TestGenerateToken
//
//	@Description: 测试生成token
//	@param t
func TestGenerateToken(t *testing.T) {
	token := &Token{
		AccessExpire: 7200,
		AccessSecret: "123",
		Payload:      nil,
	}

	jwtToken, err := token.GenerateToken()
	if err != nil {
		t.Error("Error generating token")
	} else {
		t.Log(jwtToken)
	}
}

// TestParseToken
//
//	@Description: 测试解析token
//	@param t
func TestParseToken(t *testing.T) {
	token := &Token{
		AccessExpire: 7200,
		AccessSecret: "123",
		Payload:      nil,
	}

	jwtToken, err := token.GenerateToken()
	if err != nil {
		t.Error("Error generating token")
	} else {
		t.Log(jwtToken)
	}

	parsedToken, err := token.Parse(jwtToken.Token)
	if err != nil {
		t.Error("Error parsing token", err)
	} else {
		t.Log(parsedToken)
	}
}

// TestNew
//
//	@Description: 测试创建token
//	@param t
func TestNew(t *testing.T) {
	token, err := New("123", 7200).GenerateToken()
	if err != nil {
		t.Error("Error generating token")
	} else {
		t.Log(token)
	}

	// parse
	parsedToken, err := New("123", 7200).Parse(token.Token)
	if err != nil {
		t.Error("Error parsing token", err)
	} else {
		t.Log(parsedToken.Valid)
	}
}
