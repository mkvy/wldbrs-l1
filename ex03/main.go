package main

import (
	"fmt"
	"sync"
)

func sumSquare(num int, result *int, wg *sync.WaitGroup, m *sync.Mutex) {
	//сигнализируем что горутина из группы wg завершила выполнение
	defer wg.Done()
	//избавляемся от конфликтов при записи
	m.Lock()
	*result += num * num
	m.Unlock()
}

func main() {
	//создаем переменную для синхронизации
	var wg sync.WaitGroup
	var mutex sync.Mutex
	nums := []int{2, 4, 6, 8, 10}
	//5 горутин, под кол-во элементов массива
	wg.Add(len(nums))
	res := 0
	for _, v := range nums {
		//в ходе итерации запускаем горутину функции, которая суммирует результат
		go sumSquare(v, &res, &wg, &mutex)
	}
	//ждем завершения всех горутин
	wg.Wait()
	fmt.Println("Res is: ", res)
}
