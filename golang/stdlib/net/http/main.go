package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	for _, s := range []string{
		"content-type", // Content-Type
		"ContentType",  // Contenttype
		"Contenttype",  // Contenttype
		"contentType",  // Contenttype
		"Content-Type", // Content-Type
		"Content-type", // Content-Type
	} {
		// 规范头部
		fmt.Println(http.CanonicalHeaderKey(s))
	}

	buf := bytes.NewBufferString("FF D8 ab cd FF D9") // jpg
	for _, b := range [][]byte{
		[]byte("content-type"),
		[]byte("Content-Type"),
		buf.Bytes(),
		[]byte(`{"Name": "jd"}`),
	} {
		// 检测内容类型
		fmt.Println(http.DetectContentType(b))
	}

	for _, t := range []string{
		"http/1.0",
		"HTTP/1.0",
		"http/1.1",
		"HTTP/1.1",
		"Mon Jan 2 12:00:00 2018",
	} {
		// 解析 http 版本
		fmt.Println(http.ParseHTTPVersion(t))
		fmt.Println(http.ParseTime(t))
	}

	for _, s := range []int{
		200, // OK
		301, // Moved Permanently
		302, // Found
		404, // Not Found
		500, // Internal Server Error
	} {
		// 转换
		fmt.Println(http.StatusText(s))
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 劫持
		hi, ok := w.(http.Hijacker)
		if !ok {
			log.Println("can not hijack")
		} else {
			conn, rw, err := hi.Hijack()
			if err != nil {
				log.Println(err)
			}
			defer conn.Close()

			_, err = rw.Write([]byte("Hello"))
			if err != nil {
				log.Println(err)
			}
			rw.Flush()

			s, err := rw.ReadString('\n')
			if err != nil {
				log.Println(err)
			}
			log.Printf("%v\n", s)

			return
		}

		// 限制从 request body 读取的数据大小
		limit := int64(8)
		rc := http.MaxBytesReader(w, r.Body, limit)
		defer rc.Close()
		p := make([]byte, limit)
		_, err := rc.Read(p)
		if err != nil {
			log.Printf("%+v\n", err)
		}
		log.Printf("%s\n", p)

		// 代理
		u, err := http.ProxyFromEnvironment(r)
		if err != nil {
			log.Printf("%+v\n", err)
		}
		if u != nil {
			log.Printf("%s\n", u.String())
		}

		// cookie
		http.SetCookie(w, &http.Cookie{
			Name:  "Cookie",
			Value: "session",
		})

		// 返回 404
		// http.NotFound(w, r)

		// 重定向
		// http.Redirect(w, r, "http://127.0.0.1:8890/", http.StatusMovedPermanently)

		// 返回内容
		// f, err := os.Open("main.go")
		// if err != nil {
		// 	log.Printf("%+v\n", err)
		// }
		// defer f.Close()
		// http.ServeContent(w, r, "content", time.Now(), f)

		// 返回文件
		http.ServeFile(w, r, "main.go")
	})
	log.Fatal(http.ListenAndServe(":8890", mux))
}

var _ http.RoundTripper = &transport{}

type transport struct{}

func (t *transport) RoundTrip(r *http.Request) (w *http.Response, err error) {
	return
}
