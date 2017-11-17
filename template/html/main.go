package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
	if err != nil {
		panic(err)
	}
}
