package main

import (
	"bytes"
	"testing"
)

func TestHandleUploadWithSeek(t *testing.T) {
	u := bytes.NewBufferString("")
	handleUploadWithSeek(u)
}

func TestHandleUploadWithFile(t *testing.T) {
	u := bytes.NewBufferString("")
	handleUploadWithFile(u)
}
func TestHandleUploadWithMultiReader(t *testing.T) {
	u := bytes.NewBufferString("")
	handleUploadWithMultiReader(u)
}

func TestHandleUploadWithPipe(t *testing.T) {
	u := bytes.NewBufferString("")
	handleUploadWithPipe(u)
}

func TestHandleUploadWithMultiWriter(t *testing.T) {
	u := bytes.NewBufferString("")
	handleUploadWithMultiWriter(u)
}
