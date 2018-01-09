package main

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

func BenchmarkPrint(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s", "hello.")
		fmt.Println(testing.Coverage())
	}
}

func BenchmarkTemplateParallel(b *testing.B) {
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
		}
	})
}
