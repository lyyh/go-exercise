package main

import "fmt"

func main() {
	var ch1, ch2, ch3 chan int
	var i1, i2 int
	select {
	case i1 = <-ch1:
		fmt.Printf("received ", i1, " from c1\n")
	case ch2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-ch3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	}
	ch1 <- 2
}
