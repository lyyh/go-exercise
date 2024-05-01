package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 5)
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		go func(v int) {
			wg.Add(1)
			fmt.Printf("写入 %d\n", v)
			ch <- v
		}(i)
	}
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("close")
		close(ch)
	}()

	go func() {
		for v := range ch {
			fmt.Printf("读取 %d\n", v)
			wg.Done()
		}
	}()

	wg.Wait()
}
