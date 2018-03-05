package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"log"
)

func main() {
	key := []byte("I am a key")
	for _, tc := range []string{
		"Hello",
		"I am jd",
	} {
		tcb := []byte(tc)
		mac := hmac.New(sha256.New, key)
		mac.Write(tcb)
		expectedMAC := mac.Sum(nil)
		log.Printf("%s", expectedMAC)
		// log.Println(CheckMAC(tcb, expectedMAC, key))
	}
}

// CheckMAC reports whether messageMAC is a valid HMAC tag for message.
func CheckMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
