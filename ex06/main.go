package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// через пустую структуру. аналогично можно через bool / любой тип
func doSmth1(chEnd chan struct{}) {
	for {
		select {
		case <-chEnd:
			fmt.Println("goroutine doSmth1 stops")
			time.Sleep(time.Second)
			return
		}
	}
}

// через sync.WaitGroup
func doSmth2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("goroutine doSmth2 stops")
}

// через context cancel()
func doSmth3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine doSmth3 stops")
			time.Sleep(time.Second)
			return
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// через context timeout
func doSmth4(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine doSmth4 stops")
			time.Sleep(time.Second)
			return
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}

// через закрытие канала
func doSmth5(ch chan bool) {
	for {
		// при чтении канала проверяем, закрыт ли канал
		_, opened := <-ch
		if !opened {
			fmt.Println("Goroutine doSmth5 stops")
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	//через пустую структуру. аналогично можно через bool / любой тип
	ch := make(chan struct{})
	fmt.Println("doSmth1 starts")
	go doSmth1(ch)
	ch <- struct{}{}
	close(ch)

	//через sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("doSmth2 starts")
	go doSmth2(&wg)
	wg.Wait()

	//через контекст
	ctxB := context.Background()
	ctx, cancel := context.WithCancel(ctxB)
	fmt.Println("doSmth3 starts")
	go doSmth3(ctx)
	time.Sleep(time.Second * 1)
	cancel()

	//через контекст  с таймаутом
	ctx, _ = context.WithTimeout(ctxB, 3*time.Second)
	fmt.Println("doSmth4 starts")
	go doSmth4(ctx)
	time.Sleep(4 * time.Second)

	fmt.Println("doSmth5 starts")
	ch1 := make(chan bool)
	go doSmth5(ch1)
	time.Sleep(2 * time.Second)
	close(ch1)

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("everything stopped")
	}
}
