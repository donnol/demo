package main

import "reflect"

// 分析对象的方法
func parseObjectMethod(obj interface{}) (map[string]reflect.Value, error) {
	m := make(map[string]reflect.Value)

	objType := reflect.TypeOf(obj)
	objValue := reflect.ValueOf(obj)

	for i := 0; i < objType.NumMethod(); i++ {
		method := objType.Method(i)

		objMethod := objValue.Method(i)
		m[method.Name] = objMethod
	}

	return m, nil
}
