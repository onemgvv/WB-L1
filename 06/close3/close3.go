package close3

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"
)

const workersCount = 5

func Main() {
	ctx, destroy := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	wkChan := make(chan int)

	for i := 0; i < workersCount; i++ {
		go worker(i, wkChan, ctx)
	}

	go func() {
		time.Sleep(5 * time.Second)
		destroy()
	}()

	for {
		select {
		case <-ctx.Done():
			close(wkChan)
			return
		default:
			number := rand.Intn(50) + 25
			wkChan <- number
			time.Sleep(1 * time.Second)
		}
	}
}

func worker(id int, ch <-chan int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is closed \n", id)
			return
		case data, ok := <-ch:
			if !ok {
				fmt.Print("Chan already closed \n")
				return
			}
			fmt.Printf("Data: %d (Worker %d) \n", data, id)
		}
	}
}
