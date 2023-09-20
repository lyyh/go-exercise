package main

import (
	"fmt"
)

func Foo(list *[]int) []int {
	*list = append(*list, 1, 2, 3)
	fmt.Println(cap(*list))
	return *list
}
func main() {
	// var list =
	mySlice := make([]int, 5)
	fmt.Println(cap(mySlice))
	Foo(&mySlice)
	fmt.Println(len(mySlice))
}
