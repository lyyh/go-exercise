package main

import (
	"fmt"
	"os"
)

type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type SuperMan struct {
	Human
}

func main() {
	// os.OpenFile()
	// var s SuperMan
	// // s.Human = &Human{
	// // 	Name: "",
	// // 	Age:  0,
	// // }
	// s.Age = 1
	// fmt.Printf("%+v", s.Human)
	file, err := os.OpenFile("./openFile/a.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Printf("err:%s", err.Error())
	}
	file.Write([]byte("Hello"))
	fmt.Printf("%+v", file)
	// io.Reader
}
