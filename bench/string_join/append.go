package string_join

import "time"

func Append() string {
	b := make([]byte, 0, 40)
	b = append(b, "request.ID"...)
	b = append(b, ' ')
	b = append(b, "client.Addr().String()"...)
	b = append(b, ' ')
	b = time.Now().AppendFormat(b, "2006-01-02 15:04:05.999999999 -0700 MST")
	r := string(b)
	return r
}
