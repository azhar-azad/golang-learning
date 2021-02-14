package main

import (
	"fmt"

	"github.com/azhar-azad/introducing_go/chapter_8/creating_a_package/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
