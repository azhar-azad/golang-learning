package main

import "fmt"

func main() {
	arrays()
	slices()
	maps()
	exer4()
}

func arrays() {
	fmt.Println("arrays")

	nums := [5]float64{98, 93, 77, 82, 83}

	var total float64
	for _, value := range nums {
		total += float64(value)
	}
	fmt.Println("Avg = ", total/float64(len(nums)))
}

func slices() {
	fmt.Println("slices")

	// x := make([]float64, 5) // slice of an array of 5 len, slice len is also 5
	// y := make([]float64, 5, 10) // slice of an array of 10 len, slice len is 5

	arr := [5]float64{1, 2, 3, 4, 5}
	fmt.Println("arr = ", arr)
	fmt.Println("arr[0:5] = ", arr[0:5])
	fmt.Println("arr[1:4] = ", arr[1:4])
	fmt.Println("arr[0:] = ", arr[0:]) // same as arr[0:len(arr)]
	fmt.Println("arr[:5] = ", arr[:5]) // same as arr[0:5]
	fmt.Println("arr[:] = ", arr[:])   // same as arr[0:len(arr)]

	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)

	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("slice3 = ", slice3)
	fmt.Println("slice4 = ", slice4)
}

func maps() {
	fmt.Println("maps")

	map1 := make(map[string]int)
	map1["key"] = 10
	fmt.Println(map1)

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements["Li"])

	name, isPresent := elements["Un"]
	fmt.Println(name, isPresent)

	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, isPresent := elements2["Li"]; isPresent {
		fmt.Println(el["name"], el["state"])
	}
}
