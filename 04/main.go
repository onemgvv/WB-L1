package main

import (
	"context"
	"fmt"
	"l1/04/worker"
	"math/rand"
	"os"
	"os/signal"
	"sync"
)

var workersCount = 0

// Писатель
func runWriter(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
loop:
	for {
		select {
		// ожидаем сигнал завершения от контекста
		case <-ctx.Done():
			fmt.Println("[CTRL+C]: Closing app")
			// закрываем канал
			close(ch)
			// ждем завершения всех горутин и затем выходим из цикла
			wg.Wait()
			break loop
		// Пишем в канал числа из пакета random
		default:
			ch <- int(rand.Int())
		}	
	}
}

func main() {
	fmt.Println("Введите количество воркеров: ")
	fmt.Scan(&workersCount)

	c := make(chan int)

	// Ожидаем сигнал завершения работы
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	wg := &sync.WaitGroup{}

	// добавляем в wg количество воркеров (горутин)
	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go worker.StartWorker(wg, i, c)
	}

	runWriter(ctx, wg, c)

	fmt.Println("All workers is done. App closed")
}