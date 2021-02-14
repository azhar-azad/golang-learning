package main

import "fmt"

// func main() {
// 	convertFahrToCel()
// }

func convertFahrToCel() {
	fmt.Print("Enter temp in fahrenheit: ")
	var fahr float64
	fmt.Scanf("%f", &fahr)

	cel := (fahr - 32) * 5 / 9
	fmt.Println("Celcius = ", cel)
}
