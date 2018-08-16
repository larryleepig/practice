package main

import (
	"fmt"
	"time"
)

func main() {
	go f() //like thread parallize execute
	time.Sleep(time.Microsecond * 1)
}

func f() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
