package util

import "fmt"

// Print 打印
func Print(args ...interface{}) {
	format := ""
	for i := range args {
		format += "%#v"
		if i != len(args)-1 {
			format += ", "
		}
	}
	format += "\n"
	fmt.Printf(format, args...)
}
