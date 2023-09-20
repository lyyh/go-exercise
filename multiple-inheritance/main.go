// 多重继承
package main

import (
	"fmt"
	"log"
)

type Student interface {
	Name()
}

type Person struct {
	Val string
}

func (receiver Person) Name() {
	fmt.Println(receiver.Val)
}

type Xiaomin struct {
	Val string
	*Person
	*log.Logger
}

func main() {
	var s Student = Xiaomin{
		Val: "123",
		Person: &Person{
			Val: "23234",
		},
	}
	// fmt.Printf("%+v", s)
	s.Name()
}
