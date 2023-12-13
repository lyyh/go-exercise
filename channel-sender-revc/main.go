package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Printf("receiver 2:%d\n", v)
				_ = v
			}
		}()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Printf("receiver 1:%d\n", v)
				_ = v
			}
		}()
	}

	wg.Wait()
}
