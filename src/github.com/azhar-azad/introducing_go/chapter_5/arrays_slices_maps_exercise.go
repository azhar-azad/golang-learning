package main

import "fmt"

func exer4() {
	fmt.Println("exer4")

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	min := x[0]
	for _, v := range x {
		if v < min {
			min = v
		}
	}

	fmt.Println("smallest number is ", min)
}
