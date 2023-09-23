/*
   Create: 2023/9/17
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package token_manager

// 密码算法

import (
	"github.com/golang-module/dongle"
)

// 对称算法
const (
	HEX = iota
	BASE64
	AES
	RC4
)

// 非对称加密混淆
const (
	MD5 = iota + 100
	SHA256
	SHA512
	SHAKE256
	HMAC_MD5
	HMAC_SHA256
	HMAC_SHA512
)

var cipher = dongle.NewCipher()

func init() {
	cipher.SetMode(dongle.CBC)        // CBC、CFB、OFB、CTR、ECB
	cipher.SetPadding(dongle.PKCS7)   // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
	cipher.SetKey("0123456789abcdef") // key must be 16, 24 or 32 bytes
	cipher.SetIV("0123456789abcdef")  // iv must be 16 bytes (ECB mode doesn't require setting iv)
}

var encryptMap = map[int]func(text string) string{
	HEX: func(text string) string {
		return dongle.Encode.FromString(text).ByHex().ToString()
	},
	BASE64: func(text string) string {
		return dongle.Encode.FromString(text).ByBase64().ToString()
	},
	AES: func(text string) string {
		return dongle.Encrypt.FromString(text).ByAes(cipher).ToHexString()
	},
	RC4: func(text string) string {
		return dongle.Encrypt.FromString(text).ByRc4("Apollo").ToHexString()
	},
}

var decryptMap = map[int]func(text string) string{
	HEX: func(text string) string {
		return dongle.Decode.FromString(text).ByHex().ToString()
	},
	BASE64: func(text string) string {
		return dongle.Decode.FromString(text).ByBase64().ToString()
	},
	AES: func(text string) string {
		return dongle.Decrypt.FromHexString(text).ByAes(cipher).ToString()
	},
	RC4: func(text string) string {
		return dongle.Decrypt.FromHexString(text).ByRc4("Apollo").ToString()
	},
}

var hashMap = map[int]func(text string) string{
	MD5: func(text string) string {
		return dongle.Encrypt.FromString(text).ByMd5().ToHexString()
	},
	SHA256: func(text string) string {
		return dongle.Encrypt.FromString(text).BySha256().ToHexString()
	},
	SHA512: func(text string) string {
		return dongle.Encrypt.FromString(text).BySha512().ToHexString()
	},
	SHAKE256: func(text string) string {
		return dongle.Encrypt.FromString(text).ByShake256(512).ToHexString()
	},
	HMAC_MD5: func(text string) string {
		return dongle.Encrypt.FromString(text).ByHmacMd5("Apollo").ToHexString()
	},
	HMAC_SHA256: func(text string) string {
		return dongle.Encrypt.FromString(text).ByHmacSha256("Apollo").ToHexString()
	},
	HMAC_SHA512: func(text string) string {
		return dongle.Encrypt.FromString(text).ByHmacSha512("Apollo").ToHexString()
	},
}

// Encrypt 根据类型加密
func Encrypt(m int, text string) string {
	return encryptMap[m](text)
}

// EncryptDefault 默认使用AES加密
func EncryptDefault(text string) string {
	return encryptMap[AES](text)
}

// DecryptDefault 默认使用AES解密
func DecryptDefault(text string) string {
	return decryptMap[AES](text)
}

func Decrypt(m int, text string) string {
	return decryptMap[m](text)
}

func Hash(m int, text string) string {
	return hashMap[m](text)
}
