package main

import (
	"fmt"
	"time"
)

func main() {
	// var out []*int
	// for i := 0; i < 10; i++ {
	// 	out = append(out, &i)
	// }

	// fmt.Printf("%#v", out)
	// fmt.Println(*out[0], *out[1])

	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		go func(val int) {
			fmt.Println(val)
		}(val)
	}

	time.Sleep(time.Second)
}
