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
    default:
        fmt.Println("too large")
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
}
