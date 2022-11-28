package main

import (
	"errors"
	"fmt"
)

func A(a <-chan int) {
	// for {
	select {
	case b := <-a:
		fmt.Println("A", b)
	}
	// }
}

func B(a <-chan int) {
	// for {
	select {
	case b := <-a:
		fmt.Println("B", b)
	}
	// }
}

func main() {
	err := errors.New("hhh")
	if err != nil {

	}

	ints := make(chan int)
	go B(ints)
	go A(ints)

	ints <- 1
}
