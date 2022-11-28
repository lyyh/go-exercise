package main

import (
	"fmt"
	"time"
)

func main() {
	// for i, i2 := range collection {
	//
	// }
	count := 0
	for i := range time.Tick(time.Second) {
		count++
		if count >= 5 {
			break
		}
		fmt.Println(i)
	}
}
