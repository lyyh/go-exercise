package main

import (
	"fmt"
)

// Person ...
type Person interface {
	SetAge()
	SetName()
}

// Teacher ...
type Teacher struct {
	Name string
	Age  int
}

// SetAge ...
func (p *Teacher) SetAge() {
	p.Age = 24
}

// SetName ...
func (p *Teacher) SetName() {
	p.Name = "teacher"
}

// Student  ...
type Student struct {
	Name string
	Age  int
}

// SetAge ...
func (p Student) SetAge() {
	p.Age = 16
}

// SetName ...
func (p Student) SetName() {
	p.Name = "student"
}

type Cat struct {
}

func (c *Cat) Speak() {
	fmt.Println("Speak")
}
func (c Cat) Eat() {
	fmt.Println("Eat")
}
func main() {
	var t Person = &Teacher{}
	t.SetName()

	var s Person = &Student{}
	s.SetAge()

	cat := &Cat{}
	cat.Speak()
	cat.Eat()
	var a float64
	fmt.Println(a)
}
