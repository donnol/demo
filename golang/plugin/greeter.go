package main

import "plugin"
import "fmt"
import "os"

type Greeter interface {
	Greet()
}

func main() {
	lang := "english"
	if len(os.Args) == 2 {
		lang = os.Args[1]
	}

	var mod string
	switch lang {
	case "english":
		mod = "./en/en.so"
	case "chinese":
		mod = "./ch/ch.so"
	default:
		fmt.Println("can't speak", lang)
		os.Exit(1)
	}

	plug, err := plugin.Open(mod)
	if err != nil {
		panic(err)
	}

	symGreeter, err := plug.Lookup("Greeter")
	if err != nil {
		panic(err)
	}

	symGreeter.(Greeter).Greet()
}
