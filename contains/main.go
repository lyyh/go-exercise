package main

import (
	"fmt"

	"github.com/wxnacy/wgo/arrays"
)

type People struct {
	Name string
}

func StringsContains(arr []string, val string) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

func main() {
	p := People{Name: "12312"}
	p1 := People{Name: "adsfa"}
	arr := []People{p1, p}
	arr1 := []string{"1", "2"}

	i := arrays.Contains(arr, p)
	fmt.Printf("%d", i)

	fmt.Println(StringsContains(arr1, "2"))

	fmt.Printf("%T %T", 1, int32(2))
}
