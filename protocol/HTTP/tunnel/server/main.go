package main

import (
	"io"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// handler := func(w http.ResponseWriter, req *http.Request) {
	// 	log.Printf("%#v\n", req)
	// 	if req.Method == http.MethodConnect {
	// 		urlStr := req.URL.String()
	// 		log.Printf("%#v\n", req.URL)
	// 		w.Write([]byte(urlStr))
	// 	} else {
	// 		log.Printf("%#v\n", req.URL)
	// 		w.Write([]byte("hello"))
	// 	}
	// }
	// http.HandleFunc("/", handler)
	// if err := http.ListenAndServe(":8820", nil); err != nil {
	// 	log.Fatal(err)
	// }

	// 当使用 https 时，出现错误：2018/01/18 10:29:47 server.go:2848: http: TLS handshake error from 127.0.0.1:60191: tls: oversized record received with length 20037
	// if err := http.ListenAndServeTLS(":8820", "cert.pem", "key.pem", nil); err != nil {
	// 	log.Fatal(err)
	// }

	tcpServer(":8820")
}

func tcpServer(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// 建立到目标地址的链接
	targetConn, err := net.Dial("tcp", "www.google.com.hk:443")
	if err != nil {
		log.Println(err)
		return
	}
	defer targetConn.Close()
	conn.Write([]byte("HTTP/1.0 200 Connection Established\r\n\r\n"))

	go io.Copy(targetConn, conn)
	io.Copy(conn, targetConn)
}
