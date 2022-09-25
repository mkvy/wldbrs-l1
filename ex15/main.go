package main

import "fmt"

var justString string

/*
В изначальной реализации мы возвращали срез строки justString = v[:100],
происходила утечка памяти, так как justString - глобальная переменная,
которая ссылалась на тот же слайс байтов внутри исходной строки v, причем
только первых 100 элементов из 1024,
соотвественно после завершения работы функции someFunc() память от переменной v не очищалась,
так как на нее существовала ссылка в глобальной переменной justString
*/

// для решения данной проблемы необходимо произвести копирование строки

func someFunc() {
	v := createHugeString(1 << 10)
	//выполняется глубокое копирование слайса байтов
	justString = string([]byte(v[:100]))
}

func createHugeString(strSize int) string {
	res := ""
	for i := 0; i < strSize; i++ {
		res += "a"
	}
	return res
}

func main() {
	someFunc()
	fmt.Println(justString)
}
