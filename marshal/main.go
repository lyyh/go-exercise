package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json: "name"`
	Profile
}

type Profile struct {
	Site string `json: "site"`
}

func main() {
	u := &User{
		Name: "aaa",
	}
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("%s", b)

	x := &User{}
	var s = ""
	err = json.Unmarshal([]byte(s), x)
	if err != nil {
		fmt.Println(err.Error())
	}
}
