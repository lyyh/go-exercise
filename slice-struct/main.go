package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type StuRead struct {
		Name  interface{} `json:"name"`
		Age   interface{}
		HIgh  interface{}
		sex   interface{}
		Class interface{} `json:"class"`
		Test  interface{}
	}

	//方式1：只声明，不分配内存
	var stus1 []*StuRead

	//方式2：分配初始值为0的内存
	stus2 := make([]*StuRead, 0)

	stu1 := &StuRead{"asd1", 1, 1, 1, 1, 1}
	stu2 := &StuRead{"asd2", 2, 2, 2, 2, 2}

	//由方式1和2创建的切片，都能成功追加数据
	//方式2最好分配0长度，append时会自动增长。反之指定初始长度，长度不够时不会自动增长，导致数据丢失
	stus1 = append(stus1, stu1) //因为上面stus1是切片类型的结构体指针类型，所以append的类型也必须是取的地址。
	stus2 = append(stus2, stu2) //因为上面stus2是切片类型的结构体指针类型，所以append的类型也必须是取的地址。

	//成功编码
	json1, _ := json.Marshal(stus1)
	json2, _ := json.Marshal(stus2)
	fmt.Println(string(json1))
	fmt.Println(string(json2))
}
