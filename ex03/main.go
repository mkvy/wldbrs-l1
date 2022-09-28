package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func sumSquare(num int, result *int, wg *sync.WaitGroup, m *sync.Mutex) {
	//сигнализируем что горутина из группы wg завершила выполнение
	defer wg.Done()
	//избавляемся от конфликтов при записи
	m.Lock()
	*result += num * num
	m.Unlock()
}

func sumSquare2(num int, result *uint32, wg *sync.WaitGroup) {
	//сигнализируем что горутина из группы wg завершила выполнение
	defer wg.Done()
	//используем atomic, делает операцию атомарной, более эффективно, чем mutex
	atomic.AddUint32(result, uint32(num*num))
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

	//второй вариант через атомик
	nums1 := []int{2, 4, 6, 8, 10}
	//5 горутин, под кол-во элементов массива
	wg.Add(len(nums))
	var res1 uint32
	for _, v := range nums1 {
		//в ходе итерации запускаем горутину функции, которая суммирует результат
		go sumSquare2(v, &res1, &wg)
	}
	//ждем завершения всех горутин
	wg.Wait()

	fmt.Println("Res 2 is: ", res1)
}
