package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 无缓冲
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("写入准备")
		c <- 1
		fmt.Println("写入完成")
	}()

	go func() {
		fmt.Println("准备等待")
		time.Sleep(2 * time.Second)
		fmt.Println("等待完成")
		<-c
		fmt.Println("c 读取完")
		time.Sleep(1 * time.Second)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("complete")
}
