package main

import (
	"errors"
	"fmt"
)

type CheckErr struct {
	Err error
	Msg string
}

func (c CheckErr) Error() string {
	return c.Msg
}

func main() {
	// err := &CheckErr{Msg: "123"}
	err2 := errors.New("23123")
	if _, ok := err2.(CheckErr); ok {
		fmt.Printf("123123")
	}
	// var checkErr CheckErr
	// if errors.As(err2, &checkErr) {
	// 	fmt.Printf("213123")
	// }
}
