package main

import "fmt"

var x, y, z int
var python = "good language"

func choose(i int) {
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
		fallthrough // use to execute next case
	default:
		fmt.Println("too large")
	}
}
func anotherchoose(i int) {
	switch {
	case i > 3:
		fmt.Println(i, " fuck")
	default:
		fmt.Println(i, " well")
	}
}

func main() {
	x := 34
	fmt.Println(x, y, z, python)
	// for type
	var i int = 0
	for ; i < 3; i++ {
		fmt.Println(i, "hello, golang")
	}
	// if type
	if i == 2 {
		fmt.Println("right")
	} else {
		fmt.Println("wrong")
	}
	choose(i)
	choose(2)
	anotherchoose(i)
	anotherchoose(4)
}
