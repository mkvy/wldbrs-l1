package main

import (
	"fmt"
	"sync"
)

func printSquare(num int, wg *sync.WaitGroup) {
	//сигнализируем что горутина из группы wg завершила выполнение
	defer wg.Done()
	fmt.Println(num * num)
}

func main() {
	//создаем переменную для синхронизации
	var wg sync.WaitGroup
	nums := []int{2, 4, 6, 8, 10}
	//5 горутин, под кол-во элементов массива
	wg.Add(len(nums))
	for _, v := range nums {
		//в ходе итерации запускаем горутину функции, которая печатает квадрат
		go printSquare(v, &wg)
	}
	//ждем завершения всех горутин
	wg.Wait()
}
