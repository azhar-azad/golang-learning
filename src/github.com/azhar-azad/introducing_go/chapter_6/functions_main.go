package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println("avg = ", average(xs))
	fmt.Println("total = ", add(1, 2, 3))

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("factorial of 5 = ", factorial(5))

	// defer first()
	// second()

	// defer func() {
	// 	str := recover()
	// 	fmt.Println(str)
	// }()
	// panic("PANIC")

	pointers()

	exer1()
	exer2()
	exer3()
	exer4()
	exer5()
	exer10()
	exer11()
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func pointers() {
	x := 5
	zero(&x)
	fmt.Println("x = ", x)
}

func zero(xPtr *int) {
	*xPtr = 0
}
