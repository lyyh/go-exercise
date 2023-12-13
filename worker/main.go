package main

import (
	"fmt"
)

func main() {
	workers := 3
	c := make(chan struct{})
	c1 := make(chan struct{})

	leader := func() {
		cnt := 0
		for range c {
			cnt++
			if cnt == workers {
				break
			}
		}
		fmt.Println("所有worker完成工作")
		close(c)
		close(c1)
	}

	worker := func() {
		fmt.Println("worker 完成")
		c <- struct{}{}
	}
	go leader()
	for i := 0; i < workers; i++ {
		go worker()
	}
	select {
	case _, ok := <-c1:
		if !ok {
			fmt.Print("done")
		}
	}
}
