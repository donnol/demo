// 优雅重启测试
// 	在 chrome 访问后，进程一直阻塞在 Shutdown 方法，直到 readTimeOut 或 把浏览器关闭
// 分析
// 	浏览器上 conn 有四种状态 new,active,idle和closed，shutdown关闭时会关闭idle和closed的连接
// 	在chrome浏览器中，为了复用请求提高性能，他会对后端产生，只有建立，没有发送数据的连接，也就是状态在new的连接，当shutdown触发时，这些连接是不会主动关闭的，只有readTineout触发后，new转换为closed状态后，shutdown才会关闭
// 解决
// 	设置readTimeout
// 	设置setKeepalive为false，禁止浏览器建立持久连接
// safari没有这个问题，因为它不会预先建立持久连接
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type test struct {
}

// ServeHTTP ServeHTTP
func (t *test) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
	w.Write([]byte("Hello World"))
}

func main() {
	server := &http.Server{
		Addr:        ":8585",
		Handler:     &test{},
		ReadTimeout: 10 * time.Second,
	}
	// server.RegisterOnShutdown(func() {
	// 	log.Println("long connection.")
	// 	os.Exit(0)
	// })

	go server.ListenAndServe()
	fmt.Println("running...")

	ch := make(chan os.Signal, 10)
	signal.Notify(ch, os.Kill, os.Interrupt)

	<-ch
	fmt.Println("closing...")
	// server.Shutdown(context.TODO())
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("exit!")
}
