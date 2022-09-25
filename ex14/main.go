package main

import (
	"fmt"
	"reflect"
)

// через type switch
func findType(val interface{}) {
	switch v := val.(type) {
	case int:
		fmt.Printf("Integer: %v\n", v)
	case bool:
		fmt.Printf("bool: %v\n", v)
	case string:
		fmt.Printf("String: %v\n", v)
	case chan int:
		fmt.Printf("chan int: %v\n", v)
	default:
		fmt.Printf("Not recognized\n")
	}
}

func reflectType(val interface{}) {
	fmt.Println(reflect.TypeOf(val))
}

func main() {
	//через switch type
	fmt.Println("\n----- Using switch type -----")
	var v interface{} = "s"
	findType(v)
	v = 1
	findType(v)
	v = make(chan int)
	findType(v)
	v = true
	findType(v)
	fmt.Println("\n----- Using reflect.TypeOf -----")
	//через пакет reflect
	reflectType(v)
	v = make(chan int)
	reflectType(v)
	v = 1
	reflectType(v)
	v = "s"
	reflectType(v)
}
