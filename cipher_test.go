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
	//output:Z2qARs0/+rLyV4o0LEnAVFy0otp+CQh3krVWh/hgmXJI3LCnyoWjWZO+tHrCvX7WPliAi3xv0O9G3qSWCgqR6hdmtmFpYDeubPp66i9FR5o1poOz6YJybhUW7WohVYs/R5pzv7E6pomxAoB4ifqfCJUpuZr7LIgNgb+fU0hlnDWunhE4DUQtbrR+f6Iqo4lfXq1gAZgo5WVBnHRHnf+vuTHcuKCr0WwwgNjx0KRWHjhbK7w5MzuTWJyHKu5F9uqS6zZRMtCFsNbWwNkikF/0OP/HSsrHJxshQbBHSYH6ParuN+E6c1B24/JG9r+HOrgdCFdjUU1Bv5HARLmdsUZwZb0Vk4MUmp2nKPeWbHet5oTG9VoLuLT3GmcVDJ6hHdodpj4e5AeszcLSf0sRtOSsAshYL3+MJn4nAaCmgODyIOyo+QB2C8Dn0Gve6MbD130i4jX5CyAsouvkgvjGiacH5qRydL7aBwngolc8663iVHHqotu9HEszxbSEEytlN+IAXkIySgIsHTIKWONd17qTOvvO6JfNkfPQm661LFr5cs2SO7r1ax0LNHKFEqyRDBPlWsD1C2ZWKQeRWOxXnqU/6YaRe5F6neHVTUzvF1T43G/C0OYWewfulTUgZOyiQDJabmYn8lsRUSmBx44QDLvGVH0nGscmgRDMyinazGx3lgk=
}

// TestNewCipherDecode ...
func TestNewCipherDecode(t *testing.T) {
	privateKey, e := ioutil.ReadFile("./test_key/private.pem")
	if e != nil {
		t.Error(e)
	}
	dec := NewCipherDecode(privateKey, time.Now())
	bytes, e := dec.Decode(`Z2qARs0/+rLyV4o0LEnAVFy0otp+CQh3krVWh/hgmXJI3LCnyoWjWZO+tHrCvX7WPliAi3xv0O9G3qSWCgqR6hdmtmFpYDeubPp66i9FR5o1poOz6YJybhUW7WohVYs/R5pzv7E6pomxAoB4ifqfCJUpuZr7LIgNgb+fU0hlnDWunhE4DUQtbrR+f6Iqo4lfXq1gAZgo5WVBnHRHnf+vuTHcuKCr0WwwgNjx0KRWHjhbK7w5MzuTWJyHKu5F9uqS6zZRMtCFsNbWwNkikF/0OP/HSsrHJxshQbBHSYH6ParuN+E6c1B24/JG9r+HOrgdCFdjUU1Bv5HARLmdsUZwZb0Vk4MUmp2nKPeWbHet5oTG9VoLuLT3GmcVDJ6hHdodpj4e5AeszcLSf0sRtOSsAshYL3+MJn4nAaCmgODyIOyo+QB2C8Dn0Gve6MbD130i4jX5CyAsouvkgvjGiacH5qRydL7aBwngolc8663iVHHqotu9HEszxbSEEytlN+IAXkIySgIsHTIKWONd17qTOvvO6JfNkfPQm661LFr5cs2SO7r1ax0LNHKFEqyRDBPlWsD1C2ZWKQeRWOxXnqU/6YaRe5F6neHVTUzvF1T43G/C0OYWewfulTUgZOyiQDJabmYn8lsRUSmBx44QDLvGVH0nGscmgRDMyinazGx3lgk=`)
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bytes))
	//output:/ip4/127.0.0.1/tcp/14001/ipfs/QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx
}
