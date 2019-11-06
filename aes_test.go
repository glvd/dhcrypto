package dhcrypto

import (
	"math/rand"
	"testing"
	"time"
)

// TestNewAES ...
func TestNewAES(t *testing.T) {
	tm := time.Now()
	ts := TimeToTS(tm)

	r := rand.New(rand.NewSource(tm.UnixNano()))
	n := r.Intn(2048)
	key := cipherTable[n : n+16]
	t.Log("ts", ts)
	enc := NewAESEncoder(key, tm)
	bytes, e := enc.Encode("abcdefg")
	if e != nil {
		t.Log(e)
	}
	t.Log(string(bytes))
	dec := NewAESDecode(key, tm)
	decode, e := dec.Decode(string(bytes))
	if e != nil {
		t.Log(e)
	}
	t.Log(string(decode))
}
