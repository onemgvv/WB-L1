package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	number := []int{1, 10, 4, 3, 45, 11, 12, 7, 8, 6, 5}
	chanX := make(chan int, len(number))
	chanX2 := make(chan int, len(number))

	for i := 0; i < len(number); i++ {
		go func(i int) {
			chanX <- number[i]
		}(i)

		go func(i int) {
			chanX2 <- number[i] * 2
		}(i)
	}

	for {
		select {
		case dat := <-chanX2:
			io.WriteString(os.Stdout, strconv.Itoa(dat))
			fmt.Fprint(os.Stdout, "\n")
		default:
			os.Exit(0)
		}
	}
}