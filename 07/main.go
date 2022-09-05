package main

import (
	"fmt"
	"sync"
)

func main() {
	var mx sync.Mutex
	var wg sync.WaitGroup

	dtlen := 10

	wg.Add(dtlen)

	names := []string{"Annie", "Mary", "Raf", "Jack", "Johnny", "Rey", "Ronald", "Kennie", "Khalid", "Walt"}
	numbers := []int{111, 211, 321, 241, 456, 712, 998, 999, 198, 100}

	data := make(map[string]int)

	for i := 0; i < dtlen; i++ {
		go func(i int) {
			mx.Lock()
			data[names[i]] = numbers[i]
			mx.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()
	for k, v := range data {
		fmt.Printf("Name: %s, Phone: %d\n", k, v)
	}
}