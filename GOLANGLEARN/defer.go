package main

import "fmt"

//practice defer panic recover
func test() (result int) {
	defer func() {
		result++
	}()
	return 0
}
func main() {
	fmt.Println(test())
}
