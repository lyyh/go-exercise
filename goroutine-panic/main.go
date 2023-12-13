package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("goroutine A")
		}
	}()

	go func() {
		fmt.Println("goroutine B")
		time.Sleep(time.Second)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err", err)
			}
		}()
		panic("goroutine B panic")

	}()

	time.Sleep(2 * time.Second)
	fmt.Println("main end")
	return
}
