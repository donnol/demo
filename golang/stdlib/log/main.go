package main

import (
	"bytes"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.SetOutput(os.Stdout)
	log.SetPrefix("jd's log: ")

	log.Printf("%b, %s\n", log.Flags(), log.Prefix())
	log.Printf("Hello, I am %s\n", "jd")
	// calldepth := 0 // jd's log: 2018/04/20 14:59:12 main.go:17: jd
	// calldepth := 1 // jd's log: 2018/04/20 14:59:40 log.go:351: jd
	// calldepth := 2 // jd's log: 2018/04/20 14:59:46 proc.go:198: jd
	// calldepth := 3 // jd's log: 2018/04/20 15:00:57 asm_amd64.s:2361: jd
	// calldepth := 4 // jd's log: 2018/04/20 15:01:17 ???:0: jd
	calldepth := 5 // jd's log: 2018/04/20 15:01:17 ???:0: jd
	s := "jd"
	err := log.Output(calldepth, s)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range []int{
		log.Ldate,
		log.Ltime,
		log.LstdFlags,
		log.Lmicroseconds,
		log.Llongfile,
		log.Lshortfile,
		log.LUTC,
	} {
		log.Printf("%b\n", v)
	}

	buf := new(bytes.Buffer)
	logger := log.New(buf, "new log: ", log.Lshortfile|log.LstdFlags)
	logger.Printf("Hello, I am %s\n", "jd")
	log.Printf("%s\n", buf.String())
}
