package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("sleep 2 秒")
		time.Sleep(time.Second * 2)
		wg.Done()
	}()

	go func() {
		fmt.Println("sleep 2 秒")
		time.Sleep(time.Second * 2)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("结束")
}
