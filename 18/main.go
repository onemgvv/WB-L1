package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Ticker struct {
	count int
}

func (t *Ticker) increment() {
	t.count++
}

func stop(ch chan<- os.Signal, sec time.Duration) {
	time.Sleep(5 * time.Second)
	ch <- syscall.SIGTERM
}

func main() {
	workersCount := 10
	var wg sync.WaitGroup

	t := Ticker{}

	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					t.increment()
				}
			}
		}(ctx)
		wg.Done()
	}

	quit := make(chan os.Signal, 1)

	go stop(quit, 5)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	cancel()
	wg.Wait()

	fmt.Println(t.count)
}