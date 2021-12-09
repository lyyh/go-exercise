package main

import (
	"encoding/json"
	"fmt"
)

type header struct {
	Encryption  interface{} `json:"-"`           // 结构体加标签
	Timestamp   interface{} `json:"timestamp"`   // 结构体加标签
	Key         string      `json:"key"`         // 结构体加标签
	Partnercode int         `json:"partnercode"` // 结构体加标签
	Class       *Class      `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	h := header{
		Encryption:  "sha",
		Timestamp:   1482463793,
		Key:         "2342874840784a81d4d9e335aaf76260",
		Partnercode: 10025,
	}

	//指针变量
	cla := new(Class) //这个new方法，就相当于  cla := &Class{}，是一个取地址的操作。
	cla.Name = "1班"
	cla.Grade = 3
	h.Class = cla

	b, errs := json.Marshal(h) // 序列化
	if errs != nil {
		fmt.Println(errs.Error())
	}
	// 直接打印 b 是 byte[]
	fmt.Print(string(b))
}
