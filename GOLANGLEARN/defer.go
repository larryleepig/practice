package main

import "fmt"

//practice defer panic recover
func test() (result string) {
	defer func() {
		result = "not 0"
	}()
	return "0"
}

func LIFO() {
	for i := 0; i < 4; i++ { //defer LIFO like stack
		defer fmt.Println(i)
	}
}

func main() {
	fmt.Println(test())
	LIFO()
	defer func() {
		fmt.Println("c")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("d")
	}()
	fu()
}

func fu() {
	fmt.Println("a")
	panic("hello") //panic 以下的code不會執行 中間的值會傳到recover()
	fmt.Println("b")
	fmt.Println("f")
}
