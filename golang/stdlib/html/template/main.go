package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	filename := "template.html"
	w, err := os.OpenFile(filename, os.O_TRUNC|os.O_CREATE|os.O_RDWR, 0700)
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	s := "<html></html>"
	b := []byte(s)
	template.HTMLEscape(w, b)
	es := template.HTMLEscapeString(s)
	log.Printf("%s\n", es)

	es = template.HTMLEscaper(s, "<body></body>")
	log.Printf("%s\n", es)

	for _, tc := range []interface{}{
		1,
		0,
		"1",
		"0",
		0.0,
		1.1,
		"0.0",
		"1.1",
	} {
		truth, ok := template.IsTrue(tc)
		log.Printf("ok: %v, truth: %v\n", ok, truth)
	}

	s = `alert('hello')`
	es = template.JSEscapeString(s)
	log.Printf("%s\n", es)

	es = template.URLQueryEscaper(s)
	log.Printf("%s\n", es)

	t, err := template.New("module").Parse("<html>{{.}}</html>")
	if err != nil {
		log.Fatal(err)
	}
	err = t.ExecuteTemplate(w, "module", "hello")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", t.Name())

	err = t.ExecuteTemplate(w, "module", "O'Reilly: How are <i>you</i>?")
	if err != nil {
		log.Fatal(err)
	}
}
