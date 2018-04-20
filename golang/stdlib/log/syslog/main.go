package main

import (
	"log"
	"log/syslog"
)

func main() {
	// syslogd 必须在运行
	syslogger, err := syslog.NewLogger(syslog.LOG_WARNING, log.Lshortfile|log.LstdFlags)
	if err != nil {
		log.Fatal(err) // Unix syslog delivery error
	}
	syslogger.Printf("Hello, I am %s\n", "jd")
}
