// wsl 无法使用 grmon 监控，必须使用 ubuntu 虚拟机
package main

import (
	"github.com/bcicen/grmon"
)

func main() {
	grmon.Start()

	for {
	}
}
