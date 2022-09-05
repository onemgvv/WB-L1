package main

import (
	"fmt"
	"l1/02/count_array"
)

func main() {
	fmt.Println("Count Array Sqrt")
	fmt.Println("Count using Mutex: ")
	count_array.CountArrayMutex()
	fmt.Println("Count using Channels: ")
	count_array.CountArrayChannel()
	fmt.Println("Count using Channels and WaitGroup: ")
	count_array.CountArrayChannelWg()
}