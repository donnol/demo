package main

import (
	"html"
	"log"
)

func main() {
	for _, s := range []string{
		`<html></html>`,
		"<body></body>",
		"<div></div>",
		"<div><a>a</a></div>",
	} {
		es := html.EscapeString(s)
		log.Printf("%s, %s\n", s, es)

		s = html.UnescapeString(es)
		log.Printf("%s, %s\n", s, es)
	}
}
