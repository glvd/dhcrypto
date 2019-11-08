package dhcrypto

import "C"

//NewCipherDecodeJNI ...
//export NewCipherDecodeJNI
func NewCipherDecodeJNI(key []C.byte, ts C.int64, s []C.byte) C.CString {
	println(key, ts, s)
	return ""
}
