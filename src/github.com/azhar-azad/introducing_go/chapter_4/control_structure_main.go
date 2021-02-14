package main

import "fmt"

func main() {
	loop()
	conditioning()
	exer2()
	exer3()
}

func loop() {
	fmt.Println("loop")

	i := 1
	for i <= 10 {
		fmt.Println("i = ", i)
		i++
	}

	fmt.Println()

	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			fmt.Println("j = ", j, " even")
		} else {
			fmt.Println("j = ", j, " odd")
		}
	}
}

func conditioning() {
	fmt.Println("conditioning")

	i := 3

	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	default:
		fmt.Println("Unknown")
	}
}
