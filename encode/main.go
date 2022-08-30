package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

// 定义一个结构体
type Monster struct {
	Name     string `json:"name,omitempty"`
	Age      int    `json:"-"`
	Birthday string
	Sal      float64
	Skill    string
}

func main() {
	var b = []byte{'a', 'b', 'c'}
	res := append([]byte(nil), b...)
	fmt.Println(string(res))

	m := Monster{}
	data, err := json.Marshal(&m)
	if err != nil {
		return
	}
	fmt.Println(string(data))

	var str = "1"
	fmt.Println(String2Bytes(str))

	v := [3]int{1, 2, 3}
	for _, val := range v {
		str = fmt.Sprint("%s,%s", str, val)
	}
	fmt.Println(str)
}

func String2Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
