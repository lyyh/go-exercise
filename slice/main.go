package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// str := "你好，世界！hello world"
	// s := []rune(str)
	// s[3] = '够'
	// s[4] = '浪'
	// s = s[:14]
	// str = string(s)
	// fmt.Println(str)
	// a := []string{}
	// fmt.Printf("%p", a)
	// var b []string
	// fmt.Printf("a %+v,b %+v", a, b)
	// fmt.Printf("a %d,b %d", unsafe.Sizeof(a), unsafe.Sizeof(b))
	// fmt.Print(b == nil)

	// a := getEmptySlice()
	// b := getNilSlice()
	c := make([]string, 0)
	marshal, _ := json.Marshal(c)
	fmt.Printf("%s", marshal) // null
	// for _, s := range c {
	// 	fmt.Println(s)
	// }
	// c = append(c, "1")
	// for _, s := range c {
	// 	fmt.Println(s)
	// }

	// fmt.Printf("%p", c)
	// for _, s := range c {
	// 	fmt.Print(s)
	// }
	// d := make([]string, 0)
	// fmt.Print(d == nil)
	// a = append(a, []string{"asdf"}...)
	// b = append(b, []string{"234"}...)
	// c = append(c, "a")
	// fmt.Printf("a %+v", a)
	// fmt.Printf("b %+v", b)
	// fmt.Printf("c %+v", c)
}

func getEmptySlice() []string {
	return []string{}
}

func getNilSlice() []string {
	return nil
}
