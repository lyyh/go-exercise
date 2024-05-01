package main

import (
	"fmt"
	"log"
)

type T struct {
	Name string `in:"path" name:"name"`
	*log.Logger
}

type Compose interface {
	Foo()
	Fuc()
}

func (t T) method2() {
	fmt.Println(t.Name)
}

func (t *T) method() {
	fmt.Println(t.Name)
}

func (t T) Fuc() {

}

func (t T) Foo() {

}

func Run(c Compose) {
	c.Fuc()
	c.Foo()
}

type S struct {
	T T `json:"t,omitempty"`
}

func (s *S) Format(state fmt.State, verb rune) {
	// fmt.Println(state)
	// fmt.Println(verb)
}

func (s *S) String() string {
	return "String fun" + s.T.Name
}
func main() {
	// t := T{Name: "123", Logger: log.New(os.Stdout, "", 0)}
	// // fmt.Printf("%s\n", t)
	// t.method()
	// t.Println("13123")
	// Run(t)
	// s := S{}
	// marshal, _ := json.Marshal(s)
	// s2 := string(marshal)
	// fmt.Println(s2)

	s1 := S{
		T: T{
			Name: "123",
		},
	}
	// a := fmt.Sprintf("%#v", s1)
	fmt.Printf("%v", s1)
	// var params = T(t)
	// fmt.Println(params)
}
