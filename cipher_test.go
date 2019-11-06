package crypto

import "testing"

func TestRSA_LoadPrivate(t *testing.T) {
	rsa := RSA{}
	if err := rsa.LoadPrivate("./test_key/private.pem"); err != nil {
		t.Fatal(err)
	}
	if rsa.privateKey == nil {
		t.Failed()
	}
}

func TestRSA_LoadPublic(t *testing.T) {
	rsa := RSA{}
	if err := rsa.LoadPublic("./test_key/rsa.crt"); err != nil {
		t.Fatal(err)
	}

	if rsa.publicKey == nil {
		t.Failed()
	}

}
