package main

import (
	"fmt"
	"sync"
	"time"
)

func printSquare(num int, wg *sync.WaitGroup) {
	//сигнализируем что горутина из группы wg завершила выполнение
	defer wg.Done()
	fmt.Println(num * num)
}

func printSquareBufCh(num int, ch chan struct{}) {
	fmt.Println(num * num)
	//отправляем пустую структуру в канал
	ch <- struct{}{}
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
	fmt.Println("-------")
	//буферизованный канал размером с число элементов
	ch := make(chan struct{}, len(nums))
	for _, v := range nums {
		//в ходе итерации запускаем горутину функции, которая печатает квадрат
		go printSquareBufCh(v, ch)
	}
	//считываем len(nums) раз значения полученные в канал, после чего канал блокируется
	select {
	case <-ch:
	default:
		time.Sleep(10 * time.Millisecond)
	}
}
