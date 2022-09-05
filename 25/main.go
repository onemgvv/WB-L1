package main

import (
	"context"
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func sleep2(d time.Duration) {
	tick := time.Tick(d / 100)
	for i := 0; i < 100; i++ {
		<-tick
	}
}

func sleep3(d time.Duration) {
	tmt, cancelFunc := context.WithTimeout(context.Background(), d)
	defer cancelFunc()
	<-tmt.Done()
}

func sleep4(d time.Duration) {
	tmt, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(d))
	defer cancelFunc()
	<-tmt.Done()
}

func main() {
	sleep(2 * time.Second)
	fmt.Println("(time.After) After 2 seconds")
	sleep2(2 * time.Second)
	fmt.Println("(Ticker) After 2 seconds")
	sleep3(2 * time.Second)
	fmt.Println("(context Timeout) After 2 seconds")
	sleep4(2 * time.Second)
	fmt.Println("(Context Deadline) After 2 seconds")
}