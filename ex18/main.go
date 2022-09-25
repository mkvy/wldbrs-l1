package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	cnt int
	mx  sync.Mutex
}

func InitCounter(mutex *sync.Mutex) *Counter {
	c := new(Counter)
	c.cnt = 0
	c.mx = *mutex
	return c
}

func (c *Counter) IncreaseCount() {
	//блокируем одновременный доступ к ресурсу
	c.mx.Lock()
	c.cnt++
	c.mx.Unlock()
}

func (c *Counter) GetCount() int {
	return c.cnt
}

func worker(workerId int, cntr *Counter, chEnd chan struct{}) {
	for {
		select {
		case <-chEnd:
			return
		default:
			cntr.IncreaseCount()
			fmt.Printf("WORKER %d val: %d\n", workerId, cntr.cnt)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var mutex sync.Mutex
	cntr := InitCounter(&mutex)
	ch := make(chan struct{})
	go worker(1, cntr, ch)
	go worker(2, cntr, ch)
	go worker(3, cntr, ch)
	time.Sleep(5 * time.Second)
	//отправляем сигналы на закрытие
	ch <- struct{}{}
	ch <- struct{}{}
	close(ch)
	fmt.Println("Value is", cntr.GetCount())
}