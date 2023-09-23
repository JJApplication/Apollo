/*
   Create: 2023/9/20
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package token_manager

import (
	"github.com/JJApplication/Apollo/utils"
	"github.com/golang-module/dongle"
	"testing"
	"time"
)

func init() {
	cipher.SetMode(dongle.CBC)        // CBC、CFB、OFB、CTR、ECB
	cipher.SetPadding(dongle.PKCS7)   // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
	cipher.SetKey("0123456789abcdef") // key must be 16, 24 or 32 bytes
	cipher.SetIV("0123456789abcdef")  // iv must be 16 bytes (ECB mode doesn't require setting iv)
}

func TestEncryptAndDecrypt(t *testing.T) {
	var testString = "Apollo"
	for method, fn := range encryptMap {
		t.Logf("method: %d | result: %s\n", method, fn(testString))
		t.Logf("method: %d | decode: %s\n", method, decryptMap[method](fn(testString)))
	}
}

func TestHash(t *testing.T) {
	var testString = "Apollo"
	for m, fn := range hashMap {
		t.Logf("method: %d | hash: %s\n", m, fn(testString))
	}
}

func TestGenerateToken(t *testing.T) {
	now := time.Now().Unix()
	expire := now + 60
	t.Log(utils.GetTimeString(now))
	t.Log(utils.GetTimeString(expire))
	tokenTmp := GenerateToken("127.0.0.1", now)
	// encrypt
	t.Log(tokenTmp)
	t.Log(tokenTmp.String())
	// decrypt
	tokenOrigin := DecryptToken(tokenTmp.String())
	t.Log(tokenOrigin)
}

func TestValidateToken(t *testing.T) {
	now := time.Now().Unix()
	expire := now + 60
	t.Log(utils.GetTimeString(now))
	t.Log(utils.GetTimeString(expire))
	tokenTmp := generateToken("127.0.0.1", now, 60, "Apollo")
	SetToken(tokenTmp)
	t.Log(validateValue(tokenTmp, "Apollo", "0.0.0.0", now))
	t.Log(validateValue(tokenTmp, "Apollo", "127.0.0.1", now))
}
