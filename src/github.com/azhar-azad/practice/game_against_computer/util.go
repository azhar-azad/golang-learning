package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Util struct {
}

func (util *Util) getIntegerInput(msg string) int {
	var input string
	fmt.Scanf("%s", &input)
	for {
		if util.isInteger(input) {
			i, _ := strconv.Atoi(input)
			return i
		}
		println(msg)
		fmt.Scanf("%s", &input)
	}
}

func (util *Util) isInteger(str string) bool {
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

func (util *Util) getRandomChoice() int {
	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	return randomizer.Intn(2) + 1
	// return rand.Intn(2) + 1
}

func print(text string) {
	fmt.Print(text)
}

func println(text string) {
	fmt.Println(text)
}
