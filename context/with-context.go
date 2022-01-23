package main

import (
	"context"
	"fmt"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dest := make(chan int)

		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dest <- n:
					n++
				}
			}
		}()

		return dest
	}
	// WithCancel 返回带有新 Done 通道的父节点的副本，当调用返回的 cancel 函数或当关闭父上下文的 Done 通道时，将关闭返回上下文的 Done 通道，无论先发生什么情况。
	ctx, cancel := context.WithCancel(context.Background())

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	// 取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用 cancel
	defer cancel()

}
