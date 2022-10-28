package main

import (
	"fmt"
)

type MyInt int

var a int = 1
var b MyInt = 2

func main() {
	a = int(b)
	fmt.Println(a)
}
