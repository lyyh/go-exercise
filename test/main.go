package main

import (
	"encoding/json"
	"fmt"
)

type header struct {
	Encryption  string `json:"encryption"`  // 结构体加标签
	Timestamp   int64  `json:"timestamp"`   // 结构体加标签
	Key         string `json:"key"`         // 结构体加标签
	Partnercode int    `json:"partnercode"` // 结构体加标签
}

func main() {
	h := header{
		Encryption:  "sha",
		Timestamp:   1482463793,
		Key:         "2342874840784a81d4d9e335aaf76260",
		Partnercode: 10025,
	}
	b, errs := json.Marshal(h)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Print(string(b))

	a := 123
	// b := a.(type)
	fmt.Println(b)
}
