package main

import (
	"fmt"
	"strings"
)

type test struct {
}

func checkUniqueChar(str string) bool {
	//приведем к нижнему регистру
	str = strings.ToLower(str)
	//будем использовать для проверки уникальности символов хеш таблицу, где ключ - символ строки, значение - количество вхождений в строку
	charMap := make(map[rune]int)
	for _, v := range str {
		if charMap[v] == 0 {
			charMap[v]++
		} else {
			//возвращаем false если количество вхождений больше 1
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Введите строку: ")
	var str string
	fmt.Scanf("%s", &str)
	fmt.Println("Строка состоит из неповторяющихся символов?: ", checkUniqueChar(str))
}
