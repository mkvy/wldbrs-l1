package main

import "fmt"

// первый способ, через арифметику
func arithSwap(x, y int) (int, int) {
	x = x + y
	y = x - y
	x = x - y
	return x, y
}

// второй способ, через битовые операции
func bitwSwap(x, y int) (int, int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y
	return x, y
}

func main() {
	//первый способ, через арифметику
	x := 5
	y := 10
	fmt.Printf("Before swap: %d , %d\n", x, y)
	x, y = arithSwap(x, y)
	fmt.Printf("After swap: %d , %d\n", x, y)
	//второй способ, через битовые операции
	x = 10
	y = 5
	fmt.Printf("Before swap: %d , %d\n", x, y)
	x, y = bitwSwap(x, y)
	fmt.Printf("After swap: %d , %d\n", x, y)

	//третий способ, через присвоение
	x = 10
	y = 5
	fmt.Printf("Before swap: %d , %d\n", x, y)
	x, y = y, x
	fmt.Printf("After swap: %d , %d\n", x, y)

}
