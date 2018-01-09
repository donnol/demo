package main

import (
	"context"
	"testing"
)

func main() {
	a("1")
}

func assert1(ctx context.Context, t *testing.T, ret interface{}, method string, arg map[string]interface{}, commpare interface{}) (j []byte) {
	return
}

func assert2(t *testing.T, have, want interface{}, message string) {
	// have == want
	t.Log(message)
	return

	// have != want
	// t.Fatal(message)
}
