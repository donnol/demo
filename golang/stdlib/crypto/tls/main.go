package main

import (
	"crypto/rand"
	"crypto/tls"
	"log"
	"net"
	"time"
)

func main() {
	certFile, keyFile := "cert.pem", "key.pem"
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{
		Rand:         rand.Reader,
		Certificates: []tls.Certificate{cert},
	}
	l, err := tls.Listen("tcp", ":8825", config)
	if err != nil {
		log.Fatal(err)
	}
	for {
		log.Println("before accept")
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("after accept")
		// conn = tls.Server(conn, config)
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	log.Println("handling...")
	time.Sleep(time.Second)
	_, err := conn.Write([]byte("Hello, I am jd"))
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second)
	conn.Close()
	log.Println("handle finish.")
}
