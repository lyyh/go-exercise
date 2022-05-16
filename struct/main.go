package main

import "fmt"

type T struct {
	name string
}

func (t T) method2() {
	fmt.Println(t.name)
}

func (t *T) method() {
	fmt.Println(t.name)
}

func main() {

}
