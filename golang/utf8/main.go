package main

import (
	"fmt"
	"unicode/utf8"

	sc "golang.org/x/text/encoding/simplifiedchinese"
)

func main() {
	gbkEncoder := sc.GB18030.NewDecoder()
	// r, err := gbkEncoder.Bytes([]byte{0xe6, 0xa8, 0x8a})
	word := "樊"
	fmt.Printf("%s\n", word)

	r, err := gbkEncoder.Bytes([]byte(word))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", r)

	testCase := [][]byte{
		[]byte("\b5Ὂg̀9! ℃ᾭG"),
		[]byte("我"),
		[]byte{0xff, 0xfe, 0xfd},
		[]byte(r),
		[]byte{0xB7, 0xAE, 255},
	}

	for _, s := range testCase {
		v := utf8.Valid(s)
		fmt.Println(s, v)
	}

}
