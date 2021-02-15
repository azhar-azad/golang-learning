package app

import (
	"fmt"

	"github.com/azhar-azad/practice/dps_calculator/util"
)

var packageName string

/*
	PUBLIC FUNCTIONS
*/

func AppConstructor() {
	setPackageName()
}

func GetPackageName() string {
	return packageName
}

/*
	PRIVATE FUNCTIONS
*/

func setPackageName() {
	showAvailablePackages()
	fmt.Print("Enter your choice: ")
	packageChoice := getPackageInput("Please enter integer value", 6)
	packageName = getPackageNameFromChoice(packageChoice)
}

func showAvailablePackages() {
	fmt.Println("This DPS packages are available right now.")
	fmt.Println("MMSS (press 1)")
	fmt.Println("SHWAPNO (press 2)")
	fmt.Println("SHEFA (press 3)")
	fmt.Println("FEMINA (press 4)")
	fmt.Println("SU-GRIHINI (press 5)")
}

func getPackageInput(msg string, totalPackageCount int) int {
	input := util.GetIntegerInput(msg)

	if input == 0 {
		return input
	}

	if input <= totalPackageCount {
		return input
	}

	fmt.Println("Please enter value from given choices.")
	return getPackageInput(msg, totalPackageCount)
}

func getPackageNameFromChoice(packageChoice int) string {
	switch packageChoice {
	case 0:
		return "DEMO"
	case 1:
		return "MMSS"
	case 2:
		return "SHWAPNO"
	case 3:
		return "SHEFA"
	case 4:
		return "FEMINA"
	case 5:
		return "SU-GRIHINI"
	default:
		return "INVALID"
	}
}
