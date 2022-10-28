package main

import (
	"fmt"
)

func fullName(firName *string, lastName *string) {
	defer fmt.Println("defer println")
	if firName == nil {
		s := "runtime error: first name cannot be nil"
		panic(s)
	}
	if lastName == nil {
		panic("lastName is nil")
	}
}

func main() {
	var firName string = "123"
	fullName(&firName, nil)
}
