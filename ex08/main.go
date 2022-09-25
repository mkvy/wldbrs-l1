package main

import (
	"fmt"
	"strconv"
)

func changeBit(num int64, pos uint, bitValue bool) int64 {
	//если бит = 1
	if bitValue {
		//конъюнкцией проверяем значение бита в заданной позиции
		if num&(1<<pos) == 0 {
			//если 0, возвращаем исключающее или между числом и маской
			return num ^ (1 << pos)
		}
		return num
	} else {
		//если уже 0, вернули исходное
		if num&(1<<pos) == 0 {
			return num
		}
		//иначе исключающее или
		return num ^ (1 << pos)
	}
	return num
}

func main() {
	//число
	var num int64
	num = 5
	//исходное двоичное
	fmt.Println("----- NUM -----")
	fmt.Println(strconv.FormatInt(num, 2))
	//после изменения
	var pos uint = 0
	fmt.Printf("----- AFTER CHANGING BIT WITH NUMBER %d -----\n", 3)
	fmt.Println(strconv.FormatInt(changeBit(num, pos, false), 2))
}
