package util

import "github.com/azhar-azad/practice/dps_calculator/properties"

var props properties.IProperty

func DpsUtilsConstructor(packageName string) {
	props = properties.GetProperties(packageName)
	props.Ge
}
