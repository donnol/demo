package main

import (
	"encoding/ascii85"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	data := []byte("Hello, I am dong")
	suffix := "dong"
	// suffix := "don"
	for i := 0; i < 10; i++ {
		n := len(data)
		max := ascii85.MaxEncodedLen(n) // (n + 3) / 4 * 5
		log.Printf("%d, %d\n", n, max)

		dst := make([]byte, max)
		actNum := ascii85.Encode(dst, data)
		// log.Printf("%v, %v\n", dst, dst[:actNum])
		log.Printf("%s, %s\n", dst, dst[:actNum])

		dst2 := make([]byte, n)
		ndst, nsrc, err := ascii85.Decode(dst2, dst, true)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%d, %d, %s\n", ndst, nsrc, dst2)

		data = append(data, []byte(suffix)...)
	}

	w, err := os.OpenFile("testdata", os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	wc := ascii85.NewEncoder(w)
	defer wc.Close()
	n, err := wc.Write(data) // 将 data 编码后保存到文件里
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d, %s\n", n, data)

	w.Close()

	w, err = os.OpenFile("testdata", os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	r := ascii85.NewDecoder(w)
	p := make([]byte, len(data)) // 将文件里的 码文 解码后读出来
	n, err = r.Read(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d, %s\n", n, p)

	w.Close()
}
