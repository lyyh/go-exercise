package main

import (
	"fmt"
	"strings"
)

func main() {
	a := 1
	b := string(rune(a))
	fmt.Println(b)

	lower := strings.ToLower("AAAA") // 返回一个新的字符串
	fmt.Println(lower)

	fmt.Println(strings.ToUpper("asdf"))
}
