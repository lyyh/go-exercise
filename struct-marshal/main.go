package main

import (
	"fmt"
)

type Point struct {
	X int
}
type SDXJResultJsonStr struct {
	Result []Point `json:"result"`
	State  string  `json:"state"`
}

func main() {
	s := &SDXJResultJsonStr{}
	fmt.Printf("%v", *s)
	// var str = "{}"
	// jsonStr := SDXJResultJsonStr{}
	// jsonStrByte, err := json.Marshal(jsonStr)
	// fmt.Printf("%s", string(jsonStrByte))
	// err = json.Unmarshal([]byte(str), &jsonStr)
	// if err != nil {
	// 	fmt.Printf("%s", err.Error())
	// }
	// fmt.Printf("%+v", jsonStr)
	// a := `{"result":null,"state":"\u0000"}`
	// _ = json.Unmarshal([]byte(a), &jsonStr)
	// fmt.Printf("%+v", jsonStr)
}
