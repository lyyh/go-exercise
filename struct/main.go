package main

import "fmt"

type T struct {
	Name string `in:"path" name:"name"`
}

func (t T) method2() {
	fmt.Println(t.Name)
}

func (t *T) method() {
	fmt.Println(t.Name)
}

func main() {
	t := T{name: "123"}
	var params = T(t)
	fmt.Println(params)
}
