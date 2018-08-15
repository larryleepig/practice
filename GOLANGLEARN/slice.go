package main

import "fmt"

func main() {
	var x, y []float64
	arr := [5]float64{1, 2, 3, 4, 5}
	z := arr[:3]
	x = make([]float64, 5)
	fmt.Println(x, y, z)
	y = append(z, 4, 5)
	copy(x, y)
	fmt.Println("x = ", x, "\ny = ", y, "\nz = ", z)

}
