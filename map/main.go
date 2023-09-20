package main

import (
	"fmt"
)

func main() {
	var idNameMap map[int64]string // 只定义没有初始化
	fmt.Printf("%+v\n", idNameMap)
	m := map[int64]string{} // 初始化
	fmt.Printf("%+v\n", m)
	var listSlice []string // 数组、切片会在定义的时候初始化一个默认值
	listSlice = append(listSlice, "123")
	fmt.Printf("%+v\n", listSlice)
}
