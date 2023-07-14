package stringbyteconverter

import "unsafe"

type Converter struct {
}

// New creates new converter
func New() *Converter {
	return &Converter{}
}

// ToBytes converts string to []byte
func (c *Converter) ToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// ToString converts []byte to string
func (c *Converter) ToString(b []byte) *string {
	return (*string)(unsafe.Pointer(&b))
}
