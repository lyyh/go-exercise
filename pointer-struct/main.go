package main

import (
	"fmt"
)

type Car struct {
	color string
	size  int
}

func (c *Car) name() {
	fmt.Printf("color is %s,size is %d", c.color, c.size)
}

func (c *Car) setColor(color string) {
	c.color = color
}

func (c *Car) setSize(size int) {
	c.size = size
}
func main() {
	c1 := Car{color: "#ddd", size: 100}
	c1.name()
	c1.setColor("#000")
	c1.name()
}
