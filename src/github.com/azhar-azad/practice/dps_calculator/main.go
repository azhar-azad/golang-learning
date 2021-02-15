package main

import (
	"fmt"

	"github.com/azhar-azad/practice/dps_calculator/app"
	"github.com/azhar-azad/practice/dps_calculator/calculator"
	"github.com/azhar-azad/practice/dps_calculator/table"
	"github.com/azhar-azad/practice/dps_calculator/util"
)

func main() {
	app.AppConstructor()

	packageName := app.GetPackageName()
	fmt.Println("\"" + packageName + "\" selected.\n")

	fmt.Println("1. Calculate DPS by inputting properties (press 1)")
	fmt.Println("2. Display DPS scheme in a table (press 2)")
	fmt.Print("Enter choice: ")

	choice := util.GetIntegerInput("Please enter integer value.")

	switch choice {
	case 1:
		calculator.Main(packageName)
	case 2:
		table.Main(packageName)
	default:
		fmt.Println("Invalid choice.")
	}
}
