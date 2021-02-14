package main

import "fmt"

func main() {

	numbers()
	strings()
	booleans()
}

func numbers() {
	fmt.Println("\nnumbers")
	fmt.Println("1 + 1 = ", 1+1)
}

func strings() {
	fmt.Println("\nstrings")

	fmt.Println("len(\"Hello, World\") = ", len("Hello, World"))
	fmt.Println("\"Hello, World\"[1] = ", "Hello, World"[1])
	fmt.Println("\"Hello, \" + \"World\" = ", "Hello, "+"World")
}

func booleans() {
	fmt.Println("\nbooleans")

	fmt.Println("true && true = ", true && true)
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || true = ", true || true)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
}

/*

numbers
1 + 1 =  2

strings
len("Hello, World") =  12
"Hello, World"[1] =  101
"Hello, " + "World" =  Hello, World

booleans
true && true =  true
true && false =  false
true || true =  true
true || false =  true
!true =  false

*/
