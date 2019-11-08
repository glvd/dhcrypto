package dhcrypto

import (
	"encoding/base64"
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
	tm, err := time.Parse(TimeStampFormat, "20191107")
	if err != nil {
		return
	}

	enc := NewCipherEncoder(pubKey, 4999, tm)
	bytes, e := enc.Encode("/ip4/127.0.0.1/tcp/14001/ipfs/QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx")
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bytes))
	//output:Z2qARs0/+rLyV4o0LEnAVFy0otp+CQh3krVWh/hgmXJI3LCnyoWjWZO+tHrCvX7WPliAi3xv0O9G3qSWCgqR6hdmtmFpYDeubPp66i9FR5o1poOz6YJybhUW7WohVYs/R5pzv7E6pomxAoB4ifqfCJUpuZr7LIgNgb+fU0hlnDWunhE4DUQtbrR+f6Iqo4lfXq1gAZgo5WVBnHRHnf+vuTHcuKCr0WwwgNjx0KRWHjhbK7w5MzuTWJyHKu5F9uqS6zZRMtCFsNbWwNkikF/0OP/HSsrHJxshQbBHSYH6ParuN+E6c1B24/JG9r+HOrgdCFdjUU1Bv5HARLmdsUZwZb0Vk4MUmp2nKPeWbHet5oTG9VoLuLT3GmcVDJ6hHdodpj4e5AeszcLSf0sRtOSsAshYL3+MJn4nAaCmgODyIOyo+QB2C8Dn0Gve6MbD130i4jX5CyAsouvkgvjGiacH5qRydL7aBwngolc8663iVHHqotu9HEszxbSEEytlN+IAXkIySgIsHTIKWONd17qTOvvO6JfNkfPQm661LFr5cs2SO7r1ax0LNHKFEqyRDBPlWsD1C2ZWKQeRWOxXnqU/6YaRe5F6neHVTUzvF1T43G/C0OYWewfulTUgZOyiQDJabmYn8lsRUSmBx44QDLvGVH0nGscmgRDMyinazGx3lgk=
}

// TestNewCipherDecode ...
func TestNewCipherDecode(t *testing.T) {
	privateKey, e := ioutil.ReadFile("./test_key/private.pem")
	if e != nil {
		t.Error(e)
	}
	tm, err := time.Parse(TimeStampFormat, "20191107")
	if err != nil {
		return
	}
	dec := NewCipherDecode(privateKey, tm)
	bytes, e := dec.Decode(`Z2qARs0/+rLyV4o0LEnAVFy0otp+CQh3krVWh/hgmXJI3LCnyoWjWZO+tHrCvX7WPliAi3xv0O9G3qSWCgqR6hdmtmFpYDeubPp66i9FR5o1poOz6YJybhUW7WohVYs/R5pzv7E6pomxAoB4ifqfCJUpuZr7LIgNgb+fU0hlnDWunhE4DUQtbrR+f6Iqo4lfXq1gAZgo5WVBnHRHnf+vuTHcuKCr0WwwgNjx0KRWHjhbK7w5MzuTWJyHKu5F9uqS6zZRMtCFsNbWwNkikF/0OP/HSsrHJxshQbBHSYH6ParuN+E6c1B24/JG9r+HOrgdCFdjUU1Bv5HARLmdsUZwZb0Vk4MUmp2nKPeWbHet5oTG9VoLuLT3GmcVDJ6hHdodpj4e5AeszcLSf0sRtOSsAshYL3+MJn4nAaCmgODyIOyo+QB2C8Dn0Gve6MbD130i4jX5CyAsouvkgvjGiacH5qRydL7aBwngolc8663iVHHqotu9HEszxbSEEytlN+IAXkIySgIsHTIKWONd17qTOvvO6JfNkfPQm661LFr5cs2SO7r1ax0LNHKFEqyRDBPlWsD1C2ZWKQeRWOxXnqU/6YaRe5F6neHVTUzvF1T43G/C0OYWewfulTUgZOyiQDJabmYn8lsRUSmBx44QDLvGVH0nGscmgRDMyinazGx3lgk=`)
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bytes))
	//output:/ip4/127.0.0.1/tcp/14001/ipfs/QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx
}

// TestBase64Encode ...
func TestBase64Encode(t *testing.T) {
	t.Log(string(Base64Encode([]byte("/ip4/127.0.0.1/tcp/ 14001/ \nipfs/ QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx"))))
	t.Log(base64.RawStdEncoding.EncodeToString([]byte("/ip4/127.0.0.1/tcp/ 14001/ \nipfs/ QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx")))
	t.Log(base64.URLEncoding.EncodeToString([]byte("/ip4/127.0.0.1/tcp/ 14001/ \nipfs/ QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx")))
	t.Log(base64.RawURLEncoding.EncodeToString([]byte("/ip4/127.0.0.1/tcp/ 14001/ \nipfs/ QmRnCapPN73gDHRhnmKD8VbFjsqdF6Y5adzihwbjULaTcx")))
}

// TestNewCipherDecode2 ...
func TestNewCipherDecodeETH(t *testing.T) {
	s := `IraOAjhxuuQYuTfYQsbQ90aetg/AKYO82Ro4JTe17QZkAwRoO79cob0z2RtK1/YtoNyMFzoxgTIegLflimIQmjAnP9we7vytdS3tC50Y7TlbuquUvpZsWpGTJ+dFhY2f2817zTz9aG+diHLho23UZXrEjHmeiw3+a/Zc0GSfiC4a2Y0mSiG7TEBAuY5qaQ+xC5DI15B184qEwckbmIkEG7m0uI7arcnBrBzufkqrUM6nXongVfYigauKksx4feM0N1vCtVo+WjDAZLRaXlQuPgsXCObZsI+NX7b15FxzFD/YfQKlpFUmzbRknMO1VGkH+jymkMXZE307XKMHjHD8WRJeTR6a2mSL71BMFOkDtliaLQHFIsexofNZoxrlBE4GWmwzUubdYBWcoHEiAx6G6vajzCUlRdI3D6fhbzdPV1usDXtY9+xCivqxv6x10HB7O9gL5Z9hgA8OXzXG/+YuJAivJ3DU+5urhL2J77j30P7gh9sv1HTUbsTdFHpKvkVvGFAhTgTzaU74fbPPwQA047b/weRd1aDJucg8DvN19Ftju1XnlmcFQPKg4qiX3aKhIu7iVFF5fNlQL1KOQ74W8Q/ynSByzjbSuxOhWRUpEGy/h1MNaI9bJlXMMptLHWLjTfXWjZYF8pSmD1rBUKR66gPkqd5TpulV2aqaikbJc8Q=`
	privateKey, e := ioutil.ReadFile("./test_key/private.pem")
	if e != nil {
		t.Error(e)
	}
	//tm, err := time.Parse(TimeStampFormat, "20191107")
	//if err != nil {
	//	return
	//}
	tm := time.Unix(1573205547, 0)
	dec := NewCipherDecode(privateKey, tm)
	bytes, e := dec.Decode(s)
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bytes))

}
