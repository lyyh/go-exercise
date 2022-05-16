package main

import (
	"fmt"

	"github.com/jinzhu/copier"
)

type Perm1 struct {
	Action string
	Label  string
}

type Perm2 struct {
	Action string
	Label  string
}

func main() {
	var p1 = Perm1{Action: "1", Label: "2"}
	var p2 = Perm2{}
	err := copier.Copy(&p2, &p1)
	if err == nil {
		fmt.Printf("%#v", p2)
	}

	psm1 := map[int]*[]string{1: {"a", "b", "c"}}
	psm2 := make(map[int][]string) // make(map[int][]string) 类型可以正常拷贝
	copier.Copy(&psm2, &psm1)
	fmt.Printf("%#v", psm2)
}
