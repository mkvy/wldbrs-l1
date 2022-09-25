package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func writer(ch chan<- int, strC chan struct{}) {
	for {
		select {
		case <-strC:
			return
		default:
			ch <- rand.Intn(50)
		}
	}
}

func worker(ch <-chan int, strC chan struct{}) {
	for {
		select {
		case <-strC:
			return
		default:
			fmt.Fprintf(os.Stdout, "Value recieved: %d\n", <-ch)
		}
	}
}

func main() {
	//считываем N - время по условию
	var timeC time.Duration
	fmt.Scanf("%d", &timeC)
	//структура для синхронизации
	structCh := make(chan struct{})
	//канал для передачи данных
	ch := make(chan int)
	defer close(ch)
	defer close(structCh)
	//вызываем горутины на чтение и запись данных с канала
	go worker(ch, structCh)
	go writer(ch, structCh)
	//после истечения заданного времени, посылаем в канал пустую структуру, после чтения горутины завершаются
	select {
	case <-time.After(timeC * time.Second):
		structCh <- struct{}{}
	}
}
