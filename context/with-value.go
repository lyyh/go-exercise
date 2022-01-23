package main

import (
	"context"
	"fmt"
)

func main() {
	type favContextKey string
	// 从上下文中根据key取value
	f := func(ctx context.Context, key favContextKey) {
		if v := ctx.Value(key); v != nil {
			fmt.Println("found value ", v)
			return
		}
		fmt.Println("key not found:", key)
	}

	k := favContextKey("language")
	// 创建一个携带key为k，value为"Go"的上下文
	ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)
	f(ctx, favContextKey("color"))
}
