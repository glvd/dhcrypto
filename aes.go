package dhcrypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

// AESType ...
type AESType int

// AES ...
type AES struct {
	key string
	ts  time.Time
}

func (a AES) cipherKey() string {
	return TimeToTS(a.ts) + a.key + TimeToTS(a.ts)
}

// Decode ...
func (a AES) Decode(s string) ([]byte, error) {
	key := a.cipherKey()
	ciphertext := []byte(s)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	Output("aes:decode:iv", fmt.Sprintf("%x", iv))
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)
	return PKCS7UnPadding(ciphertext), nil
}

// Encode ...
func (a *AES) Encode(s string) ([]byte, error) {
	key := a.cipherKey()
	plaintext := PKCS7Padding([]byte(s), aes.BlockSize)

	if len(plaintext)%aes.BlockSize != 0 {
		panic("plaintext is not a multiple of the block size")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// NewAESEncoder ...
func NewAESEncoder(key string, ts time.Time) Encoder {
	return &AES{key: key, ts: ts}
}

// NewAESDecode ...
func NewAESDecode(key string, ts time.Time) Decoder {
	return &AES{key: key, ts: ts}
}

/*PKCS7Padding PKCS7Padding */
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

/*PKCS7UnPadding PKCS7UnPadding */
func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if unpadding < 1 || unpadding > 32 {
		unpadding = 0
	}
	return plantText[:(length - unpadding)]
}
