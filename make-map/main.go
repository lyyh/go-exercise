package main

import "fmt"

func main() {
	// m := make(map[string]int)

	// m["k1"] = 1
	// m["k2"] = 2

	// // fmt.Printf(m)

	// v1 := m["k1"]
	// delete(m, "k2")
	// v2 := &v1
	// fmt.Println(v2)
	// fmt.Println(v1)
	// fmt.Println(len(m))

	// n := map[string]int{"a": 1}
	// fmt.Println(n)
	MapTest()
}

func MapTest() {
	// var m map[string]string
	var m = make(map[string]string)
	m["a"] = "a"
	fmt.Printf("%#v", m)
}
