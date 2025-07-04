package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// EncryptAES256 使用aes256加密
func EncryptAES256(text string, key string) (string, error) {
	keyRaw := []byte(key)
	if len(keyRaw) != 32 {
		return "", errors.New("key must be 32 bytes")
	}
	block, err := aes.NewCipher(keyRaw)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES256 使用aes256解密
func DecryptAES256(encryptText string, key string) (string, error) {
	keyRaw := []byte(key)
	if len(keyRaw) != 32 {
		return "", errors.New("key must be 32 bytes")
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encryptText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(keyRaw)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
