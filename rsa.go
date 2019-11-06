package dhcrypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// RSA ...
type RSA struct {
	//privateKey []byte
	privateKey *rsa.PrivateKey
	//publicKey []byte
	publicKey *rsa.PublicKey
}

// Decode ...
func (r *RSA) Decode(s string) ([]byte, error) {
	t, err := Base64Decode([]byte(s))
	if err != nil {
		return nil, err
	}

	b, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, r.privateKey, t, nil)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Encode ...
func (r *RSA) Encode(s string) ([]byte, error) {
	part, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, r.publicKey, []byte(s), nil)
	if err != nil {
		return nil, fmt.Errorf("EncryptOAEP error:%w", err)
	}

	buf := make([]byte, base64.StdEncoding.EncodedLen(len(part)))
	base64.StdEncoding.Encode(buf, part)
	return buf, nil
}

// LoadPublicRSAFromBytes ...
func LoadPublicRSAFromBytes(bytes []byte) (Encoder, error) {
	key, e := parseRSAPublicKeyFromPEM(bytes)
	if e != nil {
		return nil, e
	}
	return &RSA{
		privateKey: nil,
		publicKey:  key,
	}, nil
}

// LoadPublicRSAFromFile ...
func LoadPublicRSAFromFile(path string) (Encoder, error) {
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}
	return LoadPublicRSAFromBytes(bytes)
}

// LoadPrivateRSAFromBytes ...
func LoadPrivateRSAFromBytes(bytes []byte) (Decoder, error) {
	key, e := parseRSAPrivateKeyFromPEM(bytes)
	if e != nil {
		return nil, e
	}

	return &RSA{
		privateKey: key,
		publicKey:  nil,
	}, nil
}

// LoadPrivateRSAFromFile ...
func LoadPrivateRSAFromFile(path string) (Decoder, error) {
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}
	return LoadPrivateRSAFromBytes(bytes)
}

/*ParseRSAPrivateKeyFromPEM Parse PEM encoded PKCS1 or PKCS8 private key */
func parseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, ErrorKeyMustBePEMEncoded
	}

	pkey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		if parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
			var ok bool
			if pkey, ok = parsedKey.(*rsa.PrivateKey); !ok {
				return nil, ErrorNotRSAPrivateKey
			}
		}
	}
	return pkey, nil
}

/*ParseRSAPublicKeyFromPEM Parse PEM encoded PKCS1 or PKCS8 public key */
func parseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(key)
	if block == nil {
		return nil, ErrorKeyMustBePEMEncoded
	}

	// Parse the key
	parsedKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		if cert, err := x509.ParseCertificate(block.Bytes); err == nil {
			parsedKey = cert.PublicKey
		} else {
			return nil, err
		}
	}

	if pkey, ok := parsedKey.(*rsa.PublicKey); ok {
		return pkey, nil
	}

	return nil, ErrorNotRSAPublicKey

}

/*Base64Encode Base64Encode */
func Base64Encode(b []byte) []byte {
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(buf, b)
	return buf
}

/*Base64Decode Base64Decode */
func Base64Decode(b []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(b)))
	n, err := base64.StdEncoding.Decode(buf, b)
	return buf[:n], err
}

/*Base64DecodeString Base64DecodeString */
func Base64DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
