package close1

import "fmt"

func generator() chan int {
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}

func Main() {
	number := generator()
	fmt.Println(<-number)
  fmt.Println(<-number)
	close(number)
}