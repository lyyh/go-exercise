package main

import (
	"fmt"
)

type People struct {
	stuName string
}

func (p People) Name() string {
	return p.stuName
}

type Student interface {
	Name() string
}

func f(v Student) {
	// if i, ok := v.(People); ok {
	// 	fmt.Print(i)
	// }
	switch t := v.(type) {
	case People: // 非接口类型
		fmt.Printf("%T", t)
		break
	}
}
func main() {
	people := People{stuName: "asdf"}
	f(people)
}
