package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	fmt.Println(os.Args)

	execSpec := &syscall.ProcAttr{
		Env:   os.Environ(),
		Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()},
	}

	// dead loop
	// pid, err := syscall.ForkExec(os.Args[0], os.Args, execSpec)

	cmd := ""
	if len(os.Args) != 0 {
		cmd = os.Args[0]
	} else {
		// cmd = "fork"
		return
	}
	params := []string{}
	if len(os.Args) > 1 {
		params = os.Args[1:]
	}
	pid, err := syscall.ForkExec(cmd, params, execSpec)
	if err != nil {
		panic(err)
	}
	fmt.Println(pid)

	os.Exit(0)
}
