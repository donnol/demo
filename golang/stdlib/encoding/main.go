package main

import (
	"encoding"
)

type bm struct{}

func (b bm) MarshalBinary() (data []byte, err error) {
	return
}

func (b bm) UnmarshalBinary(data []byte) error {
	return nil
}

type tm struct{}

func (t tm) MarshalText() (text []byte, err error) {
	return
}

func (t tm) UnmarshalText(text []byte) error {
	return nil
}

func main() {
	var _ encoding.BinaryMarshaler = bm{}
	var _ encoding.BinaryUnmarshaler = bm{}
	var _ encoding.TextMarshaler = tm{}
	var _ encoding.TextUnmarshaler = tm{}
}
