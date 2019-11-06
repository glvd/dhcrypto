package dhcrypto

import (
	"testing"
)

func TestRSA_LoadPrivate(t *testing.T) {
	rsa := RSA{}
	if err := rsa.LoadPrivate("./test_key/private.pem"); err != nil {
		t.Fatal(err)
	}
	if rsa.privateKey == nil {
		t.Failed()
	}
	rlt, err := rsa.Decode("QkPJQZZtpfn8Y32WpvCOb1WK2eTeWy5rrWSqCU5jUFK04zW/c1IfCFDt6zJtV+GNWQsJQRCs/UObpktGcz0GGSsztNATEYAdOtCkrh6L3Ty52E2Mm/vwml50N/1hvX8zbAf0inDo+RyoyUtvGyIDAa4Fj2Gbuc/K1Ql9UNXwZVg=")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(rlt))
	//output:abcdefg
}

func TestRSA_LoadPublic(t *testing.T) {
	rsa := RSA{}
	if err := rsa.LoadPublic("./test_key/rsa.crt"); err != nil {
		t.Fatal(err)
	}

	if rsa.publicKey == nil {
		t.Failed()
	}
	rlt, err := rsa.Encode("abcdefg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(rlt))
	//output: QkPJQZZtpfn8Y32WpvCOb1WK2eTeWy5rrWSqCU5jUFK04zW/c1IfCFDt6zJtV+GNWQsJQRCs/UObpktGcz0GGSsztNATEYAdOtCkrh6L3Ty52E2Mm/vwml50N/1hvX8zbAf0inDo+RyoyUtvGyIDAa4Fj2Gbuc/K1Ql9UNXwZVg=

}
