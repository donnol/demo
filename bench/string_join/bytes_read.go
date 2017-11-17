package string_join

import (
	"bytes"
	"fmt"
)

func BytesRead() string {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s %v %v", "request.ID", "client.Addr()", "time.Now()")
	r := b.String()
	return r
}
