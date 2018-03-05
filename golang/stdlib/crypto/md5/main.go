package main

import (
	"crypto/md5"
	"log"
	"strings"
)

func main() {
	for _, tc := range []string{
		"Hello",
		"I am jd",
	} {
		tc = strings.Repeat(tc, 100)
		data := []byte(tc)

		hash := md5.New()
		_, err := hash.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		md5Data := hash.Sum(nil)
		log.Printf("%s, %d", md5Data, len(md5Data))

		md5Data2 := md5.Sum(data)
		log.Printf("%s, %d", md5Data2, len(md5Data2))
	}
}
