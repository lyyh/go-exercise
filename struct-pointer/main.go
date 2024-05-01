package main

import (
	"fmt"
)

type A struct {
	val int
}

func (a A) SetV() {
	a.val++
}
func (a *A) SetP() {
	a.val++
}
func main() {
	a := A{}
	// a.SetV()
	// fmt.Println(a.val)
	// a.SetV()
	a.SetP()
	// b := &a
	// b.SetP()
	fmt.Println(a.val)
}
