package string_join

import "fmt"

func Fmt() string {
	r := fmt.Sprintf("%s %v %v", "request.ID", "client.Addr()", "time.Now()")
	return r
}
