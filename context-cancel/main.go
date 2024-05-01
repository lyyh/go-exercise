package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx := context.Background()
	ctx, _ := context.WithTimeout(parentCtx, 1*time.Millisecond)
	// defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("f1")
		f1(ctx)
	case <-ctx.Done():
		fmt.Println("f2")
		f2(ctx)
	}
	fmt.Println("finished")
}

func f1(ctx context.Context) {
	fmt.Println("脑子进煎鱼了", ctx.Err())
}

func f2(ctx context.Context) {
	fmt.Println("煎鱼进脑子了", ctx.Err())
}
