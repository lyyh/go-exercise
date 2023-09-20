package main

import (
	"fmt"
	"log"
	"os"
)

type T struct {
	Name string `in:"path" name:"name"`
	*log.Logger
}

type Compose interface {
	Foo()
	Fuc()
}

func (t T) method2() {
	fmt.Println(t.Name)
}

func (t *T) method() {
	fmt.Println(t.Name)
}

func (t T) Fuc() {

}

func (t T) Foo() {

}

func Run(c Compose) {
	c.Fuc()
	c.Foo()
}

func main() {
	t := T{Name: "123", Logger: log.New(os.Stdout, "", 0)}
	t.method()
	t.Println("13123")
	Run(t)
	// var params = T(t)
	// fmt.Println(params)
}
