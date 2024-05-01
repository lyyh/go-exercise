package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

type AwoSome struct {
	Num  int    `json:"num"`
	Text string `json:"text"`
}

func main() {
	jsonStr := `{"Num":1,"Text":"aaaaa"}`
	some := AwoSome{}
	err := json.Unmarshal([]byte(jsonStr), &some)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", some)
	marshal, err := json.Marshal(some)
	fmt.Printf("val:%s", marshal)
	// two1 := &TwoInts{}
	// two1.a = 12
	// two1.b = 10
	// two1.Run()
	// fmt.Printf("two1 is: %v\n", two1)
	// fmt.Println("two1 is:", two1)
	// fmt.Printf("two1 is: %T\n", two1)
	// fmt.Printf("two1 is: %#v\n", two1)
}

// func (tn *TwoInts) String() string {
// 	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
// }

func (tn *TwoInts) String() string {
	return "aa (" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

func (tn *TwoInts) Run() {
	fmt.Println("TwoInts run")
}
