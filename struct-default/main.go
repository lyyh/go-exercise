package main

import (
	"fmt"
)

type Foo struct {
	Name string `json:"name" default:"123"`
	Is   bool   `default:"true"`
}

func main() {
	// f := new(Foo)
	foo := Foo{}
	fmt.Printf("%v\n", foo)
}
