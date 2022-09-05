package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func ReadTimeout() time.Duration {
	var tm time.Duration
	fmt.Println("Через сколько секунд завершаем?")
	scan, err := fmt.Scan(&tm)
	if err != nil || scan <= 0 {
		log.Fatal("Check input data")
	}

	return tm * time.Second
}

func Writer(ch chan<- int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- int(rand.Int())
		}	
	}
}

func Reader(ch <-chan int) {
	for data := range ch {
		fmt.Println("[msg]:", data)
	}
}

func main() {
	// Читаем timeout программы
	timeout := ReadTimeout()
	ch := make(chan int)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	go Reader(ch)
	Writer(ch, ctx)

	fmt.Println("End of APP")
}