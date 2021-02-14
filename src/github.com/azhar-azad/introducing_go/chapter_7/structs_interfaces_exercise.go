package main

import (
	"fmt"
	"math"
)

type Shape2 interface {
	area() float64
	perimeter() float64
}

type Circle4 struct {
	radius float64
}

func (c *Circle4) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle4) perimeter() float64 {
	return 2.0 * math.Pi * c.radius
}

type Rectangle3 struct {
	length, width float64
}

func (r *Rectangle3) area() float64 {
	return r.length * r.width
}

func (r *Rectangle3) perimeter() float64 {
	return 2.0 * (r.length + r.width)
}

func exer3() {
	fmt.Println("\nexer3")

	circle := Circle4{3.5}
	rectangle := Rectangle3{4.5, 2.5}

	fmt.Printf("circle area = %v\ncircle perimeter = %v\n", circle.area(), circle.perimeter())
	fmt.Printf("rectangle area = %v\nrectangle perimeter = %v\n", rectangle.area(), rectangle.perimeter())
}
