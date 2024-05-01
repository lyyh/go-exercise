package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type sendMsg struct {
	Name string `json:"name"`
	Info string `json:"info"`
}

func main() {
	h := json.RawMessage(`{"precomputed": true}`)

	c := struct {
		Header *json.RawMessage `json:"header"`
		Body   string           `json:"body"`
	}{Header: &h, Body: "Hello Gophers!"}

	b, err := json.MarshalIndent(&c, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	// jsonStr := `{"sendMsg":{"name":"louiswang","info":"这里是一条消息"},"say":"Hello"}`
	// var data map[string]json.RawMessage
	// if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
	// 	fmt.Println("json 解析失败了:", err)
	// 	return
	// }
	// var saystr string
	// if err := json.Unmarshal(data["say"], &saystr); err != nil {
	// 	fmt.Println("RawMessage 解析失败了:", err)
	// 	return
	//
	// }
	// fmt.Printf("%#v-------%T\n", saystr, saystr)
	// var resMsg sendMsg
	// if err := json.Unmarshal(data["sendMsg"], &resMsg); err != nil {
	// 	fmt.Println("RawMessage 解析失败了(sendMsg):", err)
	// 	return
	//
	// }
	// fmt.Printf("%#v-----%T\n", resMsg, resMsg)
	// fmt.Println(resMsg.Name, resMsg.Info)
	// json.MarshalIndent()
}
