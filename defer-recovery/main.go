package main

import (
	"fmt"
)

type Tester struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (t *Tester) Handler() {
	fmt.Printf("%#v", t)
	t.Age = 8
}

func main() {
	t := Tester{
		Name: "test",
		Age:  10,
	}
	t.Handler()
	fmt.Printf("t %+v", t)
	t1 := &Tester{
		Name: "test",
		Age:  10,
	}
	t1.Handler()
	fmt.Printf("t %+v", t1)
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Printf("aaaaa")
	// 	}
	// }()
	// go func() {
	// 	panic("123123")
	// }()
	// time.Sleep(time.Second)

}
