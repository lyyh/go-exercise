package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Run(c, i)
	}
	<-c
}

func Run(c chan bool, index int) {
	fmt.Printf("1111 <<< %d\n", index)
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}

	if index == 1 {
		c <- true
	}
}
