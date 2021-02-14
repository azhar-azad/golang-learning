package main

import (
	"fmt"
	"math"
)

type Circle2 struct {
	radius float64
}

func (c *Circle2) area() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	length, width float64
}

func (r *Rectangle) area() float64 {
	return r.length * r.width
}

func methodsDemo() {
	c2 := Circle2{5}
	fmt.Println(c2.area())

	r := Rectangle{10, 5}
	fmt.Println(r.area())
}
