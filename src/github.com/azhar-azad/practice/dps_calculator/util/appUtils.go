package util

import (
	"fmt"
	"strconv"
)

func GetIntegerInput(msg string) int {
	var input string
	fmt.Scanf("%s", &input)

	if isInteger(input) {
		i, _ := strconv.Atoi(input)
		return i
	}

	fmt.Println(msg)
	return GetIntegerInput(msg)
}

func isInteger(str string) bool {
	if str == "" {
		return false
	}

	length := len(str)
	if length == 0 {
		return false
	}

	i := 0
	for ; i < length; i++ {
		c := str[i]
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
