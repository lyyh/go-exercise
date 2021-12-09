package main

import (
	"encoding/json"
	"fmt"
)

type peerInfo struct {
	HTTPPort int `json:"http_port"`
	TCPPort  int `json:"tcp_port"`
	versiong string
}

func main() {
	var v peerInfo
	data := []byte(`{"http_port":80,"tcp_port":3306}`)
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", v)
}
