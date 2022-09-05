package sqrtsum

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var numbers = [5]int{2, 4, 6, 8, 10}

func sqrt(n int, ch chan int) {
	ch <- n * n
}

func SqrtSumMutex() {
	mx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	// задаем в wg количество горутин
	wg.Add(len(numbers))
	var sum int

	for _, val := range numbers {
		go func(mx *sync.Mutex, n int) {
			// Блокируем данные
			mx.Lock()
			sum = sum + (n*n) // some magic
			mx.Unlock() // отпускаем данные
			wg.Done()  // говорим wg что горутина завершила работу
		}(mx, val)
	}

	// ждем выполнения всех горутин
	wg.Wait()
	fmt.Println("SUM:", sum)
}


func SqrtSumAtomic() {
	wg := &sync.WaitGroup{}

	wg.Add(len(numbers))
	var sum int64

	for _, val := range numbers {
		go func(n int) {
			// используем атомарный счетчик
			// еще один  способ синхронизации данных в го
			atomic.AddInt64(&sum, int64(n * n))
			wg.Done()
		}(val)
	}

	wg.Wait()
	fmt.Println("SUM:", sum)
}

func SqrtSumChan() {
	var sum int
	ch := make(chan int)

	for _, v := range numbers {
		go sqrt(v, ch)
		sum += <- ch
	}
	close(ch)

	fmt.Println("SUM:", sum)
}