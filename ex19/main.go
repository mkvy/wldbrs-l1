package main

import "fmt"

func reverseString(s string) string {
	//т.к. символы unicode могут занимать разное количество байт в строке, приводим к рунам
	runeS := []rune(s)
	for i := 0; i < len(runeS)/2; i++ {
		runeS[i], runeS[len(runeS)-1-i] = runeS[len(runeS)-1-i], runeS[i]
	}
	return string(runeS)
}

func main() {
	//иероглифы состоят из 3-х байт, одной руны
	s1 := "杯米貓妻子"
	fmt.Printf("before reverse: %s\n", s1)
	fmt.Printf("after reverse: %s\n", reverseString(s1))

	//в данной строке первый символ состоит из одного байта(одной руны), второй из 12 байтов(четырех рун),
	//третий из одной руны, 1 байта,
	//четвертый из одной руны, 2 байт
	s2 := "éक्षिaф"
	fmt.Printf("before reverse: %s\n", s2)
	fmt.Printf("after reverse: %s\n", reverseString(s2))
}
