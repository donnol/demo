package main

import (
	"fmt"
	"net/http"
	"reflect"
)

var mux *http.ServeMux

func main() {
	mux = http.NewServeMux()
	initHandler("/login", &Login{})
	initHandler("/user", &User{})
	s := &http.Server{
		Addr:    ":9009",
		Handler: mux,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func initHandler(pattern string, obj interface{}) {
	methodMap, err := parseObjectMethod(obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(methodMap)

	// for name, method := range methodMap {
	// 如果这样取method的话，method一直会是最后一个值，因为method是拷贝

	for name, _ := range methodMap {
		// 解决办法，在for循环中建立局部变量，将值拷贝出来
		method := methodMap[name]

		mux.HandleFunc(pattern+"/"+name, func(resp http.ResponseWriter, req *http.Request) {

			// name会一直打印循环的最后一个键
			// method会是一个指针值
			fmt.Println(name, method)

			// 业务
			result := method.Call([]reflect.Value{reflect.ValueOf(resp), reflect.ValueOf(req)})

			// 返回
			resp.Write([]byte(fmt.Sprintf("%+v\n", result)))
		})
	}
}
