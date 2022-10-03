package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func worker(i int, ch <-chan int) {
	for d := range ch {
		//получаем число и пишем в поток
		fmt.Fprintf(os.Stdout, "Worker: %d    | Value recieved: %d\n", i, d)
	}
}

//воркеры завершаются при завершении работы main, после приема сигнала прерывания и остановки посыла данных в канал

func main() {
	//считываем N воркеров
	var wCount int
	fmt.Scanf("%d", &wCount)
	//канал для передачи данных
	ch := make(chan int)
	for i := 0; i < wCount; i++ {
		//запускаем воркеров
		go worker(i+1, ch)
	}

	//определяем канал для обработки прерывания
	done := make(chan os.Signal, 1)
	//принимаем сигнал прерывания, ctrl-c, завершения процесса
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	//в бесконечном цикле ждем, когда поступит сигнал прерывания, в случае если поступил, закрываем канал и возвращаемся из main
	for {
		select {
		//обработка прерывания
		case <-done:
			close(ch)
			fmt.Fprintln(os.Stdout, "Program exited, interrupted")
			return
		default:
			ch <- rand.Intn(50)
		}
	}
}
