package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(time.Second * 10)

	stop <- true

	time.Sleep(time.Second * 5)
}
