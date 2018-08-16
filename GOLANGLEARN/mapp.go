package main

import "fmt"

func main() {
	//var x map[string]int /this is wrong
	//this is right
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x)
	fmt.Println(x["key"])

	y := make(map[int]int)
	y[1] = 10
	y[2] = 23
	fmt.Println(y)
	fmt.Println(y[1])

	delete(y, 1)
	fmt.Println(y)

	a, b := y[2]
	fmt.Println(a, b)
	if a, b := y[2]; b {
		fmt.Println("test:", a, b)
	}

	elements := map[string]map[string]string{
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

	if el, ok := elements["B"]; ok {
		fmt.Println(el["name"], el["state"])
	}

}
