package main

import (
	"fmt"
)

type s struct{}

func (s s) Write(b []byte) (n int, err error) {
	return
}

func (s s) Width() (wid int, ok bool) {
	return
}
func (s s) Precision() (prec int, ok bool) {
	return
}
func (s s) Flag(c int) bool {
	return false
}

type f struct{}

func (f f) Format(s fmt.State, c rune) {

}

type str struct{}

func (s str) String() string {
	return ""
}

type gostr struct{}

func (g gostr) GoString() string {
	return ""
}

type ss struct{}

func (s ss) ReadRune() (r rune, size int, err error) {
	return
}

func (s ss) UnreadRune() error {
	return nil
}

func (s ss) SkipSpace() {}

func (s ss) Token(skipSpace bool, f func(rune) bool) (token []byte, err error) {
	return
}

func (s ss) Width() (wid int, ok bool) {
	return
}

func (s ss) Read(buf []byte) (n int, err error) {
	return
}

type sr struct{}

func (s sr) Scan(state fmt.ScanState, verb rune) error {
	return nil
}

func main() {
	var _ fmt.State = s{}
	var _ fmt.Formatter = f{}
	var _ fmt.Stringer = str{}
	var _ fmt.GoStringer = gostr{}
	var _ fmt.ScanState = ss{}
	var _ fmt.Scanner = sr{}
}
