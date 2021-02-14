package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle2 struct {
	length, width float64
}

func (r *Rectangle2) area() float64 {
	return r.length * r.width
}

type Circle3 struct {
	radius float64
}

func (c *Circle3) area() float64 {
	return math.Pi * c.radius * c.radius
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func interfacesDemo() {
	c3 := Circle3{5}
	fmt.Println(c3.area())

	r := Rectangle2{10, 5}
	fmt.Println(r.area())

	fmt.Println("Total area = ", totalArea(&c3, &r))

	multiShape := MultiShape{
		shapes: []Shape{
			&Circle3{5},
			&Rectangle2{10, 5},
		},
	}

	fmt.Println("Multishape total area = ", multiShape.area())
}

func totalArea(shapes ...Shape) float64 {
	var totalArea float64
	for _, s := range shapes {
		totalArea += s.area()
	}
	return totalArea
}
