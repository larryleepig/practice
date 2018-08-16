package main

import (
	"fmt"
	"time"
)

func main() {
	go f() //like thread parallize execute
	time.Sleep(time.Microsecond * 1)
	s := make(chan string) // 宣告一個 channel var

	go func() {
		s <- "hello"
	}() // write in channel (sender)
	val := <-s // read channel (receiver)
	fmt.Println(val)

	t := make(chan string, 1)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("sender hello", i)
			t <- fmt.Sprintf("receiver hello %d", i)
		}
	}()

	for i := 0; i < 3; i++ {
		val = <-t
		fmt.Println(val)
	}
	//實作sellect
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		fmt.Println("c2 sender")
		c2 <- "i an the two"
	}()

	go func() {
		fmt.Println("c1 sender")
		c1 <- "i am the one"
	}()

	for i := 0; i < 3; i++ {
		select { // use to choose which chan in
		case <-c1:
			fmt.Println("fk")
		case <-c2:
			fmt.Println("cow")
		default:
			fmt.Println("nothing!")
		}
	}
	//實作timeout
	timeout := make(chan bool)

	go func() {
		time.Sleep(time.Second)
		timeout <- true
	}()

	ch := make(chan int)
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("\n\ntimeout!!!")
	}
}

func f() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}
