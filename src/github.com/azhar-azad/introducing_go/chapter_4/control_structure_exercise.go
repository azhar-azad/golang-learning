package main

import "fmt"

func exer2() {
	fmt.Println("exer2")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func exer3() {
	fmt.Println("exer3")

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, " FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, " Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, " Buzz")
		} else {

		}
	}
}
