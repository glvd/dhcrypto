package dhcrypto

import "errors"

/* all defined errors */
var (
	ErrorKeyMustBePEMEncoded = errors.New("key must be pem encoded")
	ErrorNotECPublicKey      = errors.New("key is not a valid ECDSA public key")
	ErrorNotECPrivateKey     = errors.New("key is not a valid ECDSA private key")
	ErrorNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
	ErrorNotRSAPublicKey     = errors.New("key is not a valid RSA public key")
)
