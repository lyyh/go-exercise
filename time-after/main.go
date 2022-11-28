package main

import (
	"fmt"
	"time"
)

func pythonSender(ch chan string) {
	time.Sleep(100 * time.Second)
	ch <- "Python Sender"
}
func golangSender(ch chan string) {
	time.Sleep(500 * time.Second)
	ch <- "Golang Sender"
}
func main() {
	// fmt.Println("嗨客网(www.haicoder.net)")
	// chStr1 := make(chan string)
	// chStr2 := make(chan string)
	// go pythonSender(chStr1)
	// go golangSender(chStr2)
	// // 使用 break 可以结束 select
	// select {
	// case str1 := <-chStr1:
	// 	fmt.Println(str1)
	// case str2 := <-chStr2:
	// 	fmt.Println(str2)
	// case <-time.After(10 * time.Second):
	// 	fmt.Println("Timed out")
	// 	break
	// }

	// t := time.NewTicker(time.Millisecond * 500)
	//
	// for {
	// 	select {
	// 	// case <-t.C:
	// 	// 	fmt.Println("t.C 11111")
	// 	// 	break
	// 	case <-time.After(5 * time.Second):
	// 		fmt.Println("asdfasdf")
	// 		t.Stop()
	// 		return
	// 	}
	// }

	i := 0
Loop:
	for {
		select {
		case <-time.After(time.Second * time.Duration(2)):
			i++
			if i == 5 {
				fmt.Println("跳出for循环")
				break Loop
			}
		}
		fmt.Println("for循环内 i=", i)
	}

	fmt.Println("for循环外")
}
