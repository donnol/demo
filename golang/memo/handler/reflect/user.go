package main

import "net/http"

type User struct {
	Name string
}

func (u *User) One(resp http.ResponseWriter, req *http.Request) (User, error) {
	return User{"jd"}, nil
}

func (u *User) List(resp http.ResponseWriter, req *http.Request) ([]User, error) {
	return []User{{"jd"}, {"jc"}}, nil
}
