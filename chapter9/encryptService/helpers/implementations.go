package helpers

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type EncryptServiceInstance struct{}

var initVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}
var errEmpty = errors.New("secret key or text should not be empty")

func (EncryptServiceInstance) Encrypt(_ context.Context, key string, text string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (EncryptServiceInstance) Decrypt(_ context.Context, key string, text string) (string, error) {
	if key == "" || text == "" {
		return "", errEmpty
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	cipherText, _ := base64.StdEncoding.DecodeString(text)
	cfb := cipher.NewCFBEncrypter(block, initVector)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
