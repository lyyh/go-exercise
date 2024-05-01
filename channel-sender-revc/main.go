package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(a int) {
			fmt.Println(a)
			wg.Done()
		}(i)
	}
	wg.Wait()

	// c := make(chan int, 10)
	// wg := sync.WaitGroup{}
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// 	fmt.Println("close done")
	// }()
	//
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for v := range c {
	// 			fmt.Printf("receiver 2:%d\n", v)
	// 			_ = v
	// 		}
	// 		fmt.Println("receiver 2 out range")
	// 	}()
	// }
	//
	// for i := 0; i < 10; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		for v := range c {
	// 			fmt.Printf("receiver 1:%d\n", v)
	// 			_ = v
	// 		}
	// 		fmt.Println("receiver 1 out range")
	// 	}()
	// }
	//
	// wg.Wait()
}
