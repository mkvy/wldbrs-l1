package main

import (
	"fmt"
	"sync"
)

// пишем данные с массива в канал
func writeArray(data []int, ch chan<- int) {
	for _, v := range data {
		ch <- v
	}
	//закрываем после записи
	close(ch)
}

func readAndMultiply(chRead <-chan int, chWrite chan<- int) {
	for val := range chRead {
		chWrite <- val * 2
	}
	//закрываем после записи
	close(chWrite)
}

func main() {
	//начальный массив
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	inCh := make(chan int)
	outCh := make(chan int)
	go writeArray(arr, inCh)
	go readAndMultiply(inCh, outCh)

	//считываем, для синхронизации waitgroup
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range outCh {
			fmt.Println("value x*2 = ", val)
		}
	}()
	wg.Wait()

}
