package main

import "fmt"

func main() {
	s := make([]int, 0, 4)
	s = append(s, []int{1, 2, 3, 4}...)
	fmt.Println(len(s), " s ", cap(s))
	a := append(s, 5)
	fmt.Println(a)
	fmt.Println(len(a), " a ", cap(a))
	fmt.Println(len(s), " s ", cap(s))
	b := append(s, 6)
	fmt.Println(len(b), " b ", cap(b))
	fmt.Println(len(a), " a ", cap(a))
	fmt.Println(len(s), " s ", cap(s))
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(s)

	fmt.Println(s[4])
}
