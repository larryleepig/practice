package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (s, t, z int) {
	s = sum / 2
	t = sum - s
	return s, t, z
}
func main() {
	fmt.Println(add(43, 44))
	a, b := swap("average", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
