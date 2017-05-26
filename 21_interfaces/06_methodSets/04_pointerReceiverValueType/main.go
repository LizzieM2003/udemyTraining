package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area is", s.area())
}

func main() {
	var c circle
	c = circle{5}
	info(c)
}
