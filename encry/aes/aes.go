package aes

import (
	"bytes"
	cryptoAes "crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/forgoer/openssl"
)

var _ Aes = (*aes)(nil)

// “The length of the key can be 16/24/32 characters (128/192/256 bits)”，
// 这个key的长度只能是16,24和32个字符，分别对应AES-128, AES-192, or AES-256等模式。

// Aes aes
type Aes interface {
	i()
	// Encrypt CBC加密
	Encrypt(encryptStr string) (string, error)

	// Decrypt CBC解密
	Decrypt(decryptStr string) (string, error)

	// ECBEncrypt ECB加密
	ECBEncrypt(encryptStr string) (string, error)

	// ECBDecrypt ECB解密
	ECBDecrypt(decryptStr string) (string, error)

	// Des3CBCEncrypt 3DES-CBC 加密
	// 实现 https://github.com/forgoer/openssl
	Des3CBCEncrypt(encryptStr string) (string, error)
	// Des3CBCDecrypt 3DES-CBC 解密
	Des3CBCDecrypt(decryptStr string) (string, error)
}

type aes struct {
	key string
	iv  string
}

func New(key, iv string) Aes {
	return &aes{
		key: key,
		iv:  iv,
	}
}

func (a *aes) i() {}

func (a *aes) ECBEncrypt(encryptStr string) (string, error) {
	encryptBytes := []byte(encryptStr)
	block, _ := cryptoAes.NewCipher([]byte(a.key))
	encryptBytes = pkcs5Padding(encryptBytes, block.BlockSize())
	encrypted := make([]byte, len(encryptBytes))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(encryptBytes); bs, be = bs+size, be+size {
		block.Encrypt(encrypted[bs:be], encryptBytes[bs:be])
	}

	return base64.URLEncoding.EncodeToString(encrypted), nil
}

func (a *aes) ECBDecrypt(decryptStr string) (string, error) {
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}
	block, _ := cryptoAes.NewCipher([]byte(a.key))
	decrypted := make([]byte, len(decryptBytes))
	size := block.BlockSize()

	for bs, be := 0, size; bs < len(decryptBytes); bs, be = bs+size, be+size {
		block.Decrypt(decrypted[bs:be], decryptBytes[bs:be])
	}

	return string(pkcs5UnPadding(decrypted)), nil
}

func (a *aes) Encrypt(encryptStr string) (string, error) {
	encryptBytes := []byte(encryptStr)
	block, err := cryptoAes.NewCipher([]byte(a.key))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encryptBytes = pkcs5Padding(encryptBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(a.iv))
	encrypted := make([]byte, len(encryptBytes))
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return base64.URLEncoding.EncodeToString(encrypted), nil
}

func (a *aes) Decrypt(decryptStr string) (string, error) {
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}

	block, err := cryptoAes.NewCipher([]byte(a.key))
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(a.iv))
	decrypted := make([]byte, len(decryptBytes))

	blockMode.CryptBlocks(decrypted, decryptBytes)
	decrypted = pkcs5UnPadding(decrypted)
	return string(decrypted), nil
}

func (a *aes) Des3CBCEncrypt(encryptStr string) (string, error) {
	rv, err := openssl.Des3CBCEncrypt([]byte(encryptStr), []byte(a.key), []byte(a.iv), openssl.PKCS7_PADDING)

	return base64.RawURLEncoding.EncodeToString(rv), err
}

func (a *aes) Des3CBCDecrypt(decryptStr string) (string, error) {
	decryptBytes, err := base64.RawURLEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}
	rv, err := openssl.Des3CBCDecrypt(decryptBytes, []byte(a.key), []byte(a.iv), openssl.PKCS7_PADDING)
	return string(rv), err
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs5UnPadding(decrypted []byte) []byte {
	length := len(decrypted)
	unPadding := int(decrypted[length-1])
	return decrypted[:(length - unPadding)]
}
