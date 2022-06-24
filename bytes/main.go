package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

func main() {
	// var str = "123,234,345"
	// var bytes = [][]byte(str)
	s := [][]byte{[]byte("123"), []byte("3212")}
	fmt.Printf("%s", bytes.Join(s, []byte(",")))
	TestBytes()
}

func TestBytes() {
	f, err := ioutil.ReadFile("./tmp")
	if err != nil {
		fmt.Println("read fail:", err)
	}

	fmt.Printf("%s", f)
}
