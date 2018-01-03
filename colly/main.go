package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.CacheDir = "./_cache/"
	c.UserAgent = "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"

	// 2.找出页面里的所有连接
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// 请求连接
		e.Request.Visit(e.Attr("href"))
	})

	// 3.请求连接时
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// 4.连接返回时
	c.OnResponse(func(w *colly.Response) {
		fmt.Println("Response", string(w.Body))
	})

	c.Visit("https://wenku.baidu.com/view/197534bd011ca300a7c39019") // 1.从这个页面开始
}
