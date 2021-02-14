package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func embeddedTypesDemo() {

	p := Person{"azad"}
	p.Talk()

	a := new(Android)
	a.Talk()
}
