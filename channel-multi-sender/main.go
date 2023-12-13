package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int, 10)
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.TODO())

	go func() {
		time.Sleep(time.Second * 5)
		close(c)
		cancel()
	}()

	for id := 0; id < 10; id++ {
		go func(id int, ctx context.Context) {
			select {
			case <-ctx.Done():
				return
			case c <- id:
				fmt.Printf("入队id: %d \n", id)
			// 入队
			default:
				fmt.Printf("aaa")
			}
		}(id, ctx)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				fmt.Printf("消费id: %d\n", v)
				_ = v
			}
		}()
	}

	wg.Wait()
}
