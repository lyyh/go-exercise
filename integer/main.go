package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := 3.14
	fmt.Printf("%s", unsafe.Sizeof(i))
	fmt.Printf("%d", int(i))
}
