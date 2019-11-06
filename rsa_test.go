package dhcrypto

import (
	"testing"
)

func TestRSA_LoadPrivate(t *testing.T) {
	dec, err := LoadPrivateRSAFromFile("./test_key/private.pem")
	if err != nil {
		t.Fatal(err)
	}
	rlt, err := dec.Decode("QkPJQZZtpfn8Y32WpvCOb1WK2eTeWy5rrWSqCU5jUFK04zW/c1IfCFDt6zJtV+GNWQsJQRCs/UObpktGcz0GGSsztNATEYAdOtCkrh6L3Ty52E2Mm/vwml50N/1hvX8zbAf0inDo+RyoyUtvGyIDAa4Fj2Gbuc/K1Ql9UNXwZVg=")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(rlt))
	//output:abcdefg
}

func TestRSA_LoadPublic(t *testing.T) {
	enc, err := LoadPublicRSAFromFile("./test_key/rsa.crt")
	if err != nil {
		t.Fatal(err)
	}

	rlt, err := enc.Encode("abcdefg")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(rlt))
	//output: QkPJQZZtpfn8Y32WpvCOb1WK2eTeWy5rrWSqCU5jUFK04zW/c1IfCFDt6zJtV+GNWQsJQRCs/UObpktGcz0GGSsztNATEYAdOtCkrh6L3Ty52E2Mm/vwml50N/1hvX8zbAf0inDo+RyoyUtvGyIDAa4Fj2Gbuc/K1Ql9UNXwZVg=

}
