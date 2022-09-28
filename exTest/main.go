package main

import (
	"fmt"
)

func gor(ch chan int) {
	fmt.Println("WAITING FOR DATA")
	for z := range ch {
		fmt.Println(z)
	}
	fmt.Println("rel gor")
}

func main() {

	ch := make(chan int, 3)
	defer close(ch)

	ch <- 1
	ch <- 1
	ch <- 1
	gor(ch)
	ch <- 1
	ch <- 1

	fmt.Println("end")
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("not now")
}
