package main

import (
	"fmt"
	"net/http"
)

type Login struct {
}

func (l *Login) Login(resp http.ResponseWriter, req *http.Request) (string, error) {
	return fmt.Sprintf("hello, i am jd. req is %+v", req), nil
}

func (l *Login) Logout(resp http.ResponseWriter, req *http.Request) error {
	fmt.Println("i am logout")
	return nil
}
