package main

import (
	"fmt"
)

func main() {
	s := []string{"1", "2"}
	// s1 := s[3:]
	// fmt.Println(s1)
	for _, v := range s[1:] {
		fmt.Println(v)
	}
	fmt.Println(s)
	// a := 1
	// b := 2
	// c := float64(b-a) / float64(10) / 60
	// fmt.Println(c)
	// tasks := []int{1, 2, 3}
	// group := sync.WaitGroup{}
	// for i, _ := range tasks {
	// 	group.Add(1)
	// 	go func(v int) {
	// 		ticker := time.NewTicker(5 * time.Second)
	// 		for range ticker.C {
	// 			fmt.Printf("===== %d \n", v)
	// 			time.Sleep(20 * time.Second)
	// 		}
	// 	}(i)
	// }
	// group.Wait()
}
