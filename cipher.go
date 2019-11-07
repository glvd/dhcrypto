package dhcrypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/goextension/tool"
)

var debug = false

// TimeStampFormat ...
const TimeStampFormat = "20060102"

var cipherTable = `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`

type cipherJSON struct {
	Data      []byte
	Index     int
	Timestamp int64
}

// Decode ...
func (c cipherJSON) Decode(s string) ([]byte, error) {
	byteS, e := Base64DecodeString(s)
	if e != nil {
		return nil, e
	}
	e = json.Unmarshal(byteS, &c)
	if e != nil {
		return nil, e
	}
	Output("cipher:decode:json", string(c.Data), c.Index, c.Timestamp)
	t := time.Unix(c.Timestamp, 0)
	Output("cipher:decode:ts", t.Format(TimeStampFormat))
	key := GetCipher(c.Index, 16)
	Output("cipher:decode:key", key)
	dec := NewAESDecode(key, t)
	bytes, e := dec.Decode(string(c.Data))
	if e != nil {
		return nil, e
	}
	return bytes, nil
}

// Encode ...
func (c cipherJSON) Encode(s string) ([]byte, error) {
	t := time.Unix(c.Timestamp, 0)
	Output("cipher:encode:ts", t.Format(TimeStampFormat))
	key := GetCipher(c.Index, 16)
	Output("cipher:encode:key", key)
	enc := NewAESEncoder(key, t)
	bytes, e := enc.Encode(s)
	if e != nil {
		return nil, e
	}
	c.Data = bytes
	Output("cipher:decode:json", string(c.Data), c.Index, c.Timestamp)
	marshaled, e := json.Marshal(c)
	if e != nil {
		return nil, e
	}
	return Base64Encode(marshaled), nil
}

// NewCipherEncoder ...
func NewCipherEncoder(i int, t time.Time) Encoder {
	return cipherJSON{Index: i, Timestamp: t.Unix()}
}

// NewCipherDecode ...
func NewCipherDecode() Decoder {
	return cipherJSON{}
}

// Encoder ...
type Encoder interface {
	Encode(s string) ([]byte, error)
}

// Decoder ...
type Decoder interface {
	Decode(s string) ([]byte, error)
}

// MakeTable ...
func MakeTable() string {
	cipherTable = tool.GenerateRandomString(2048)
	return cipherTable
}

// LoadTable ...
func LoadTable(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	cipherTable = string(bytes)
	return nil
}

// SaveTable ...
func SaveTable(path string) error {
	return ioutil.WriteFile(path, []byte(cipherTable), 0755)
}

// TimeToTS ...
func TimeToTS(t time.Time) string {
	return t.Format(TimeStampFormat)
}

// GetCipher ...
func GetCipher(index, limit int) string {
	size := len(cipherTable)
	index = index % 2048
	if index+limit >= size {
		return cipherTable[index:size] + GetCipher(0, index+limit-size)
	}
	return cipherTable[index : index+limit]
}

// Output ...
func Output(v ...interface{}) {
	if debug {
		fmt.Println(v...)
	}
}
