package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID  int64 `json:"id"`
	Age int64 `json:"age"`
}

func FindTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		json := t.Field(i).Tag.Get("json")
		fmt.Println("json:", json)
	}
}
func main() {
	user := User{
		ID:  0,
		Age: 0,
	}
	FindTag(&user)
	// userType := reflect.TypeOf(user)
	// for _, i := range userType.NumField() {
	//
	// }
}
