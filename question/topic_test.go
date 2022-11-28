package main

import (
	"fmt"
	"testing"
)

type P struct {
}
type T struct {
}

var (
	m  map[P]T
	m1 = make(map[P]T)
	m2 = map[string]int{}
)

func TestTopic(t *testing.T) {
	a := (-3) % 2
	b := (-3) % (-2)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("%#v", m)
	m[P{}] = T{}
}

func TestNewStruct(t *testing.T) {
	p := new(P) // 返回指针
	p2 := &P{}
	fmt.Printf("%#v %#v", p, p2)
}

type Apple struct {
	Size int
}

type Fruit struct {
	Apple
	fruitType string
}

func TestNestStruct(t *testing.T) {
	fruit := Fruit{
		Apple:     Apple{},
		fruitType: "",
	}
	fmt.Printf("%#v", fruit)
}
