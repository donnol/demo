package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestMain(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 启动服务器
	go main()

	time.Sleep(time.Second)

	httpAddr := os.Getenv("ADDR")
	addr := "http://localhost:" + httpAddr

	getFunc := func() {
		get, err := http.NewRequest(http.MethodGet, addr, nil)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := http.DefaultClient.Do(get)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(string(content))
	}
	_ = getFunc

	// 开启多个 goroutine 不停写入
	client := http.Client{
		Timeout: time.Second * 5,
	}
	var wg sync.WaitGroup
	for i := 0; i < 300; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data := []byte(`{"BPM": 5}`)
			buf := bytes.NewReader(data)
			req, err := http.NewRequest(http.MethodPost, addr, buf)
			if err != nil {
				t.Fatal(err)
			}
			_, err = client.Do(req)
			if err != nil {
				t.Fatal(err)
			}
		}()
	}
	wg.Wait()

	// 查看结果
	// getFunc()
}
