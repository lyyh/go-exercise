package main

import (
	"fmt"
	"time"
)

func main() {
	a := float64(time.Now().UnixMilli()) / 1e3
	b := float64(time.Now().Unix())
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%f %d", a, b)
}
