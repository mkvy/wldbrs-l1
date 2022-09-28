package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type CounterInterface interface {
	IncreaseCount()
}

type CounterMutex struct {
	cnt int
	mx  sync.Mutex
}

func InitCounter() *CounterMutex {
	c := new(CounterMutex)
	c.cnt = 0
	return c
}

func (c *CounterMutex) IncreaseCount() {
	//блокируем одновременный доступ к ресурсу
	c.mx.Lock()
	c.cnt++
	c.mx.Unlock()
}

type CounterAtomic struct {
	cnt uint32
}

func (c *CounterAtomic) IncreaseCount() {
	//второй вариант, увеличиваем счетчик через atomic
	atomic.AddUint32(&c.cnt, 1)
}

func (c *CounterMutex) GetCount() int {
	c.mx.Lock()
	defer c.mx.Unlock()
	return c.cnt
}

func (c *CounterAtomic) GetCount() uint32 {
	return c.cnt
}

func worker(workerId int, cntr CounterInterface, chEnd chan struct{}) {
	for {
		select {
		case <-chEnd:
			return
		default:
			cntr.IncreaseCount()
			//fmt.Printf("WORKER %d\n", workerId)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	cntr := InitCounter()
	ch := make(chan struct{})
	go worker(1, cntr, ch)
	go worker(2, cntr, ch)
	go worker(3, cntr, ch)
	time.Sleep(5 * time.Second)
	//отправляем сигналы на закрытие
	ch <- struct{}{}
	ch <- struct{}{}
	ch <- struct{}{}
	close(ch)
	fmt.Println("---------------------------------------")
	fmt.Println("Value is", cntr.GetCount())
	fmt.Println("---------------------------------------")
	time.Sleep(5 * time.Second)
	cntrAtomic := new(CounterAtomic)
	ch1 := make(chan struct{})
	go worker(1, cntrAtomic, ch1)
	go worker(2, cntrAtomic, ch1)
	go worker(3, cntrAtomic, ch1)
	time.Sleep(5 * time.Second)
	//отправляем сигналы на закрытие
	ch1 <- struct{}{}
	ch1 <- struct{}{}
	ch1 <- struct{}{}
	close(ch1)
	fmt.Println("---------------------------------------")
	fmt.Println("Value atomic is", cntr.GetCount())
	fmt.Println("---------------------------------------")

}
