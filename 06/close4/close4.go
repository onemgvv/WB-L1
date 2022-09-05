package close4

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func Worker(ch <-chan int) {
	for data := range ch {
		fmt.Println(data)
	}
}

func Writer(ch chan<- int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- rand.Int()
		}
	}
}

func Main() {
	timeout := 2 * time.Second
	ch := make(chan int)

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	go Worker(ch)
	Writer(ch, ctx)

	fmt.Println("Closed with deadline")
}