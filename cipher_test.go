package dhcrypto

import (
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
	enc := NewCipherEncoder(4999, time.Now())
	bytes, e := enc.Encode("abdefg")
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bytes))
	//output:eyJEYXRhIjoiNkRnQjI2WE9Eb01OK0FjSUgwOUZ2UkZCT3NMT1J3eGFXbjVnVkVZd2pVbz0iLCJJbmRleCI6NDk5OSwiVGltZXN0YW1wIjoxNTczMTA1MzM2fQ==
}

// TestNewCipherDecode ...
func TestNewCipherDecode(t *testing.T) {
	dec := NewCipherDecode()
	bytes, e := dec.Decode(`eyJEYXRhIjoiNkRnQjI2WE9Eb01OK0FjSUgwOUZ2UkZCT3NMT1J3eGFXbjVnVkVZd2pVbz0iLCJJbmRleCI6NDk5OSwiVGltZXN0YW1wIjoxNTczMTA1MzM2fQ==`)
	if e != nil {
		t.Error(e)
	}
	t.Log(string(bytes))
	//output:abdefg
}
