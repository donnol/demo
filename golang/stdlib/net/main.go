package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// var
	log.Printf("%#v\n", net.IPv4bcast)

	// Conn
	var _ net.Conn = myConn{}

	// Dial
	conn, err := net.Dial("tcp", "www.jdscript.com:80") // 建立连接
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")           // 发送请求，不发送请求会一直阻塞
	status, err := bufio.NewReader(conn).ReadString('\n') // 读取返回
	if err != nil {
		log.Fatal(err)
	}
	log.Println(status)

	// 超时
	net.DialTimeout("tcp", "google.com:80", time.Second)

	// FileConn
	// f, err := os.OpenFile("file.sock", os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeSocket)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// fConn, err := net.FileConn(f)
	// if err != nil {
	// 	log.Fatal(err) // file file+net about.md: getsockopt: socket operation on non-socket
	// }
	// _ = fConn
	// b := make([]byte, 20)
	// fn, err := fConn.Read(b)
	// if err != nil && err != io.EOF {
	// 	log.Fatal(err)
	// }
	// defer fConn.Close()
	// log.Printf("%d, %s\n", fn, b)

	// Pipe
	rc, wc := net.Pipe() // 往 wc 写的时候，rc 就会读到数据
	data := "Hello, I am jd"
	go func() {
		rcb := make([]byte, len(data))
		rcn, err := rc.Read(rcb)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(rcn, string(rcb))
	}()
	wcn, err := wc.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(wcn)

	addrs, err := net.InterfaceAddrs() // 系统单播接口地址列表
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addrs)
	inters, err := net.Interfaces() // 系统网络接口列表
	if err != nil {
		log.Fatal(err)
	}
	log.Println(inters)

	log.Println(net.JoinHostPort("127.0.0.1", "8080"))  // 127.0.0.1:8080
	log.Println(net.JoinHostPort("127.0.0.1:", "8080")) // [127.0.0.1:]:8080
	log.Println(net.JoinHostPort("127.0.0.1", ":8080")) // 127.0.0.1::8080
	log.Println(net.SplitHostPort("127.0.0.1:8080"))

	for i, addr := range []string{
		"127.0.0.1",      // []
		"127.0.0.1:8080", // lookup 127.0.0.1:8080: unrecognized address
		"www.jdscript.com",
		"http://www.jdscript.com",
		"47.52.195.21", // lookup 47.52.195.21: dnsquery: DNS name does not exist.
	} {
		names, err := net.LookupAddr(addr)
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "names:", names)

		cname, err := net.LookupCNAME(addr) // canonical name 规范名称
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "cname:", cname)

		addrs, err := net.LookupHost(addr) // 主机地址
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "host:", addrs)

		ip, err := net.LookupIP(addr) // ip 地址
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "ip:", ip)

		mx, err := net.LookupMX(addr) // DNS MX 记录
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "mx:", mx)

		ns, err := net.LookupNS(addr) // DNS NS 记录
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "ns:", ns)

		port, err := net.LookupPort(addr, "http") // 端口
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "port:", port)

		cnames, srcAddrs, err := net.LookupSRV("http", "tcp", addr) // ?
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "cnames:", cnames, "addrs:", srcAddrs)

		txt, err := net.LookupTXT(addr) // DNS txt 记录
		if err != nil {
			log.Println(err)
		}
		log.Println(i, "txt:", txt)
	}

	// buf := new(net.Buffers)
	buf := net.Buffers([][]byte{[]byte("Hello, I am jd")})

	// 读出来后，内容就没了
	// w := bytes.NewBufferString("Hello, I am jd")
	// n, err := buf.WriteTo(w)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(n, w.String())
	// log.Printf("%d, %v\n", len(buf), buf)

	p := make([]byte, len([]byte("Hello, I am jd")))
	n1, err := buf.Read(p)
	if err != nil && err != io.EOF { // 读完会报 EOF 错误
		log.Fatal(err)
	}
	log.Println(n1, string(p))
}

type myConn struct{}

func (m myConn) Read(b []byte) (n int, err error) {
	return
}
func (m myConn) Write(b []byte) (n int, err error) {
	return
}
func (m myConn) Close() (err error) {
	return
}
func (m myConn) LocalAddr() (a net.Addr) {
	return
}
func (m myConn) RemoteAddr() (a net.Addr) {
	return
}
func (m myConn) SetDeadline(t time.Time) (err error) {
	return
}
func (m myConn) SetReadDeadline(t time.Time) (err error) {
	return
}
func (m myConn) SetWriteDeadline(t time.Time) (err error) {
	return
}

type myAddr struct{}

func (m myAddr) Network() string {
	return ""
}
func (m myAddr) String() string {
	return ""
}

var _ net.Listener = myListener{}

type myListener struct{}

func (m myListener) Accept() (c net.Conn, err error) {
	return
}
func (m myListener) Close() (err error) {
	return
}
func (m myListener) Addr() (a net.Addr) {
	return
}
