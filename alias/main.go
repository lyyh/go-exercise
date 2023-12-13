package main

import (
	"fmt"
)

type Name struct {
}
type KK interface {
}

func main() {
	type myString = string
	str := "A"
	a := myString(str)
	b := myString(str + "B")
	fmt.Printf("str1 type is %T\n", a)
	fmt.Printf("str2 type is %T\n", b)
	fmt.Printf("str == a is %t\n", str == a)
	fmt.Printf("a > b is %t\n", a > b)

	c := []byte{'1', '2'}
	d := string(c)
	fmt.Printf("d type is %T\n", d)

	e := "123"
	f := []byte(e)
	fmt.Printf("f type is %T,value is %v\n", f, f) // []unit8

	g := '1'
	fmt.Printf("g type is %T\n", g) // int32

	var str1 []myString
	fmt.Printf("%v\n", str1 == nil)

	var name *Name
	fmt.Printf("name %t\n", name == nil)

	var k KK
	fmt.Printf("k == nil is %t", k == nil)
}
