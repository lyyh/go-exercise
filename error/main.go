package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("a.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	if err, ok := err.(*os.PathError); ok {
		fmt.Println(err.Op)
	}
	fmt.Println(file)
}
