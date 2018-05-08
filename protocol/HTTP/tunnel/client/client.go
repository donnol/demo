package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conn, err := net.Dial("tcp", "www.jdscript.com:8820")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	farendAddress := "www.google.com.hk:443"
	req := fmt.Sprintf("CONNECT %s HTTP/1.1\r\nHost: %s\r\n\r\n", farendAddress, farendAddress)
	// req := fmt.Sprintf("GET / HTTP/1.1\r\nHost: %s\r\n\r\n", "localhost")
	log.Println(req)
	_, err = conn.Write([]byte(req))
	if err != nil {
		log.Fatal(err)
	}

	readBuf := bufio.NewReader(conn)
	for {
		respLine, err := readBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
			}
			log.Fatal(err)
		}
		log.Println(respLine)
	}
}
