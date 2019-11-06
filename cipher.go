package crypto

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)

type RSA struct {
	//privateKey []byte
	privateKey *rsa.PrivateKey
	//publicKey []byte
	publicKey *rsa.PublicKey
}

func (rsa *RSA) LoadPublic(path string) error {
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	key, e := parseRSAPublicKeyFromPEM(bytes)
	if e != nil {
		return e
	}
	rsa.publicKey = key
	return nil
}

func (rsa *RSA) LoadPrivate(path string) error {
	bytes, e := ioutil.ReadFile(path)
	if e != nil {
		return e
	}
	key, e := parseRSAPrivateKeyFromPEM(bytes)
	if e != nil {
		return e
	}
	rsa.privateKey = key
	return nil
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
