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
