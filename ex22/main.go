package main

//используем пакет big для работы с длинной арифметикой
import (
	"fmt"
	"math/big"
)

type BigNums struct {
	a *big.Int
	b *big.Int
}

func InitBigNums(a, b string) *BigNums {
	bn := new(BigNums)
	bn.a = new(big.Int)
	bn.b = new(big.Int)
	bn.a.SetString(a, 10)
	bn.b.SetString(b, 10)
	return bn
}

func (bn *BigNums) sum() string {
	sum := new(big.Int)
	sum.Add(bn.a, bn.b)
	return sum.Text(10)
}

func (bn *BigNums) multiply() string {
	prod := new(big.Int)
	prod.Mul(bn.a, bn.b)
	return prod.Text(10)
}

func (bn *BigNums) divAB() string {
	qtnt := new(big.Int)
	qtnt.Div(bn.a, bn.b)
	return qtnt.Text(10)
}

func (bn *BigNums) divBA() string {
	qtnt := new(big.Int)
	qtnt.Div(bn.b, bn.a)
	return qtnt.Text(10)
}

func (bn *BigNums) modAB() string {
	qtnt := new(big.Int)
	qtnt.Mod(bn.a, bn.b)
	return qtnt.Text(10)
}

func (bn *BigNums) modBA() string {
	qtnt := new(big.Int)
	qtnt.Mod(bn.b, bn.a)
	return qtnt.Text(10)
}

func (bn *BigNums) subAB() string {
	diff := new(big.Int)
	diff.Sub(bn.a, bn.b)
	return diff.Text(10)
}

func (bn *BigNums) subBA() string {
	diff := new(big.Int)
	diff.Sub(bn.b, bn.a)
	return diff.Text(10)
}

func main() {
	//задаем большие числа в строках, передаем в функцию конструктор
	var a string = "100000000000000000000000000000000000"
	var b string = "5555555555555555555555"
	fmt.Printf("a = 		 %s\nb = 		 %s\n", a, b)
	bn := InitBigNums(a, b)
	fmt.Println("----------------------------------------")
	fmt.Println("a + b = 	", bn.sum())
	fmt.Println("a * b = 	", bn.multiply())
	fmt.Println("a / b = 	", bn.divAB())
	fmt.Println("b / a = 	", bn.divBA())
	fmt.Println("a % b = 	", bn.modAB())
	fmt.Println("b % a= 		", bn.modBA())
	fmt.Println("a - b = 	", bn.subAB())
	fmt.Println("b - a = 	", bn.subBA())

}
