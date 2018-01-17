package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	testURL := &url.URL{
		Scheme: "https",
		Host:   "localhost:8820",
	}
	req := &http.Request{
		Method: http.MethodConnect,
		URL:    testURL,
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 因为证书是自己签发的，需要设置忽略证书校验
		},
	}
	client := http.Client{Transport: tr}
	resp, err := client.Do(req)
	// resp, err := tr.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(content))
}
