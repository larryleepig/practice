package main

import "fmt"

func main() {
	/*var x [4]float64
	for i := 0; i < 4; i++ {
		x[i] = float64(i * 27)
	}*/
	x := [4]float64{
		23,
		45,
		33,
		21,
	}

	var total float64 = 0
	for _, s := range x {
		total += s
	}
	fmt.Println(total / float64(len(x)))
}
