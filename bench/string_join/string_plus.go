package string_join

func StringPlus() string {
	s := "request.ID"
	s += " " + "client.Addr().String()"
	s += " " + "time.Now().String()"
	r := s
	return r
}
