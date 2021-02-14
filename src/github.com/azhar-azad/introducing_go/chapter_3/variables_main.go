package main

import (
	"fmt"
)

func main() {
	variables()
	convertFahrToCel()
}

func variables() {
	var x string = "hello"
	var y = "world"
	z := "!"
	fmt.Println("x == y = ", x == y)
	fmt.Println("z", z)

	var (
		a = 5
		b = 10
		c = 15
	)
	fmt.Println(a, b, c)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println("output = ", output)
}

/*

x == y =  false
z !
5 10 15
Enter a number: 4
output =  8
Enter temp in fahrenheit: 98.65
Celcius =  37.02777777777778

*/
