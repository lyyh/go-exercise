package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// var s = ""
	// id, _ := strconv.Atoi(s)
	// fmt.Printf("%d", id)
	jsonResult := json.RawMessage{}
	json.Unmarshal([]byte(`{"result":[{"tag":"fire","score":0.87,"bndbox":
[52,648,253,952]}]}`), &jsonResult)
	fmt.Printf("11 %+v", jsonResult)
}
