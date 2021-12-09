package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("1号完成")
		wg.Done()
	}()
	// wg.Done()

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("2号完成")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("好了，大家都干完了，放工")
}
