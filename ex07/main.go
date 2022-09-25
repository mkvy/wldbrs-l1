package main

import (
	"fmt"
	"sync"
)

func writeMap(data map[int]int, id int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	data[id] = id * id
	fmt.Printf("goroutine %d has written to map\n", id)
	m.Unlock()
}

func main() {
	//начальное количество данных
	n := 5
	datamap := make(map[int]int)
	//mutex для блокировки доступа к мапе
	var mutex sync.Mutex
	//WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 1; i < n+1; i++ {
		go writeMap(datamap, i, &mutex, &wg)
	}
	wg.Wait()
	for i, v := range datamap {
		fmt.Printf("id: %d ; data: %d\n", i, v)
	}
}
