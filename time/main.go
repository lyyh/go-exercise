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
	fmt.Printf("%f %d\n", a, b)
	now := time.Now()
	c := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	fmt.Println(c.Unix())

	var s *time.Time
	fmt.Println(*s)
}
