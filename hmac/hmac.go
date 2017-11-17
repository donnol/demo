package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	testCase := []struct {
		Name string
		Data []byte
	}{
		{"1", []byte("hello")},
	}

	key := []byte("haha")
	for _, single := range testCase {
		r := HMAC(single.Data, key)
		fmt.Printf("%v\n", r)

		fmt.Println(hmac.Equal([]byte{29, 62, 165, 226, 127, 183, 137, 37, 30, 240, 13, 93, 20, 167, 76, 134, 95, 112, 65, 209, 62, 179, 119, 239, 132, 53, 244, 136, 84, 250, 144, 58}, r))
	}
}

func HMAC(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	_, err := mac.Write(message)
	if err != nil {
		panic(err)
	}
	return mac.Sum(nil)
}
