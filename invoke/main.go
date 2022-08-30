package main

import (
	"fmt"
)

func Test(a, b int) int {
	defer fmt.Print("aaaa")
	defer fmt.Print("bbbb")
	return a + b
}

func main() {
	ret := Test(1, 2)
	fmt.Println(ret)
}
