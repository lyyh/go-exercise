package main

import "fmt"

func recv(ch chan int) {
	res := <-ch
	fmt.Println("res", res)
}

func main() {
	ch := make(chan int)
	go recv(ch) // 启用 goroutine 从通道接收值
	ch <- 10
	fmt.Println("发送成功！")
}
