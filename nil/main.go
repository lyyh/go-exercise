package main

import "fmt"

func main() {
	var intPointer *int
	var a interface{} = intPointer
	var b interface{}

	fmt.Println(intPointer == nil)
	fmt.Println(a == nil)
	fmt.Println(b == nil)
}
