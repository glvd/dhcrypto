package dhcrypto

import (
	"io/ioutil"
	"testing"
	"time"
)

func init() {
	debug = true
}

// TestMakeTable ...
func TestMakeTable(t *testing.T) {
	t.Log(MakeTable())
}

// TestSaveTable ...
func TestSaveTable(t *testing.T) {
	MakeTable()
	t.Log(SaveTable("ciphertable.txt"))
}

// TestGetCipher ...
func TestGetCipher(t *testing.T) {
	for i := 3162; i < 3499; i++ {
		t.Log(GetCipher(i, 16))
	}
}

// TestNewCipherEncoder ...
func TestNewCipherEncoder(t *testing.T) {
	pubKey, e := ioutil.ReadFile("./test_key/rsa.crt")
	if e != nil {
		t.Error(e)
	}
	enc := NewCipherEncoder(pubKey, 4999, time.Now())
	bytes, e := enc.Encode("/ip4/127.0.0.1/tcp/14001/ipfs/QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx")
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bytes))
	//output:RFYKjBNSC05/ahQx4HrH4d065avhVoAZHwTMPHRtnjcSb8vZA8L50q2vVXyuqr01iPdgUvRMON1Pwy9VEXVhKVhL4TgZQKEHVeBXHUK1hFphnPwIzHyID+CCAoI7opOzjSNBwTGEJc2bWPa0UTBo3yyURkE9oDJtUzegaRJjMGoVYdzXnv4/3DtqRa7cDWtucmroioLRUe5MTIbds9go4UAxY+329Cc/7Rwal+0eEjFLLNME6oEnZspio6DdGG1eGenrke5nOoHCs54JXXG2QsQpTJvj2AsA0hNGp5N8wzpIiQ7/lkqSS5hh4Oez6zHljW9cl4sF5npz7cA9hMYFBb46Ad+FiIWcihKysXjvUe0BE+BIn6dJck6PWyoKuLbUb/Rk1TtYhzCyggWeoDPI44mhCQG/1m4aZyNEPwzaB3w02H/tbCRxjEJP76r8646X8GMbPYQ/JgC+vD/eb9y4P3cSOxiwbAzRfdRWlNxkCwm44/LwRz64TIb+/EtPGeMu1/GXtmfPebXKnTQURRWoxv7V05YCU0NqEs7eA3jVCKChOzEniglOH1GmEyLWQ/+TQSSnHyNFhUnIOKnkSGaLC0NYp/RzhbvfX1rVG4f5xo5aBt6AmHDqT4oeUXIeQQh8ggDjo6pvboHHhto05e0DXjvkr1mcDQCxcWUQLem88H8=
}

// TestNewCipherDecode ...
func TestNewCipherDecode(t *testing.T) {
	privateKey, e := ioutil.ReadFile("./test_key/private.pem")
	if e != nil {
		t.Error(e)
	}
	dec := NewCipherDecode(privateKey)
	bytes, e := dec.Decode(`RFYKjBNSC05/ahQx4HrH4d065avhVoAZHwTMPHRtnjcSb8vZA8L50q2vVXyuqr01iPdgUvRMON1Pwy9VEXVhKVhL4TgZQKEHVeBXHUK1hFphnPwIzHyID+CCAoI7opOzjSNBwTGEJc2bWPa0UTBo3yyURkE9oDJtUzegaRJjMGoVYdzXnv4/3DtqRa7cDWtucmroioLRUe5MTIbds9go4UAxY+329Cc/7Rwal+0eEjFLLNME6oEnZspio6DdGG1eGenrke5nOoHCs54JXXG2QsQpTJvj2AsA0hNGp5N8wzpIiQ7/lkqSS5hh4Oez6zHljW9cl4sF5npz7cA9hMYFBb46Ad+FiIWcihKysXjvUe0BE+BIn6dJck6PWyoKuLbUb/Rk1TtYhzCyggWeoDPI44mhCQG/1m4aZyNEPwzaB3w02H/tbCRxjEJP76r8646X8GMbPYQ/JgC+vD/eb9y4P3cSOxiwbAzRfdRWlNxkCwm44/LwRz64TIb+/EtPGeMu1/GXtmfPebXKnTQURRWoxv7V05YCU0NqEs7eA3jVCKChOzEniglOH1GmEyLWQ/+TQSSnHyNFhUnIOKnkSGaLC0NYp/RzhbvfX1rVG4f5xo5aBt6AmHDqT4oeUXIeQQh8ggDjo6pvboHHhto05e0DXjvkr1mcDQCxcWUQLem88H8=`)
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bytes))
	//output:/ip4/127.0.0.1/tcp/14001/ipfs/QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx
}
