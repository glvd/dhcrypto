//go:generate go build -buildmode=c-shared -o crypto.so ./jni
package main

import "C"
import (
	"time"

	"dhcrypto"
)

//NewCipherDecodeJNI ...
//export NewCipherDecodeJNI
func NewCipherDecodeJNI(key *C.char, ts C.long, s *C.char) *C.char {
	println(key, ts, s)
	gokey := C.GoString(key)
	tm := time.Unix(int64(ts), 0)
	dec := C.GoString(s)

	decoder := dhcrypto.NewCipherDecode([]byte(gokey), tm)
	bytes, e := decoder.Decode(dec)
	if e != nil {
		return nil
	}
	str := string(bytes)
	return C.CString(str)
}

func main() {
	println("jni")
}
