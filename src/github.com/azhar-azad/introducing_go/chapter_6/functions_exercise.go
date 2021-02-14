package main

import "fmt"

func exer1() {
	fmt.Println("\nexer1")

	sliceX := []float64{1, 2, 3, 4}
	fmt.Println("sum of {1, 2, 3, 4} = ", sum(sliceX))
}

func exer2() {
	fmt.Println("\nexer2")

	half, isEven := halfOf(2)
	fmt.Printf("half(2) = (%v, %v)\n", half, isEven)
}

func exer3() {
	fmt.Println("\nexer3")

	fmt.Println("Greatest number = ", getGreatest(5, 8, 3, -3))
}

func exer4() {
	fmt.Println("\nexer4")

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}

func exer5() {
	fmt.Println("\nexer5")

	var n int
	fmt.Scanf("%v", &n)
	fmt.Println("Fib(n) = ", fib(n))
}

func exer10() {
	fmt.Println("\nexer10")

	x := 1.5
	fmt.Println("&x = ", &x)
	square(&x)
	fmt.Println("x = ", x)
}

func exer11() {
	fmt.Println("\nexer11")

	a, b := 5, 10
	fmt.Println("a, b = ", a, b)
	a, b = swap(a, b)
	fmt.Println("a, b = ", a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}

func square(x *float64) {
	*x = *x * *x
}

func fib(n int) int {
	if n == 1 {
		return 1
	}
	return n * fib(n-1)
}

func makeOddGenerator() func() uint {
	num := uint(1)
	return func() uint {
		temp := num
		num += 2
		return temp
	}
}

func getGreatest(args ...int) int {
	max := 0
	for _, num := range args {
		if num > max {
			max = num
		}
	}
	return max
}

func halfOf(num int) (int, bool) {
	return num / 2, num%2 == 0
}

func sum(nums []float64) float64 {
	total := 0.0
	for _, v := range nums {
		total += v
	}
	return total
}
