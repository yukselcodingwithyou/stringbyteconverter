package stringbyteconverter

import "unsafe"

type ToByteConverter interface {
	Convert(s string) []byte
}

type ToStringConverter interface {
	Convert(b []byte) *string
}

type toByteConverter struct{}

func (t *toByteConverter) Convert(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

type toStringConverter struct{}

func (t *toStringConverter) Convert(b []byte) *string {
	return (*string)(unsafe.Pointer(&b))
}

func NewToByteConverter() ToByteConverter {
	return &toByteConverter{}
}

func NewToStringConverter() ToStringConverter {
	return &toStringConverter{}
}
