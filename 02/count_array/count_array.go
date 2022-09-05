package count_array

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

var numbers = [5]int{2, 4, 6, 8, 10}

func CountArrayMutex() {
	mx := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	// устанавливаем количество горутин для wg.Wait
	wg.Add(len(numbers))
	for _, val := range numbers {
		go func(mx *sync.Mutex, num int) {
			// Блокируем запись в os.Stdout
			// Благодаря блокировке в момент времени доступ к данным имеет только одна горутина
			// Остальные горутины блокируются до mx.Unlock
			// Это обеспечивает целостность общих ресурсов
			mx.Lock()
			os.Stdout.WriteString((strconv.Itoa(num * num) + "\n"))
			mx.Unlock()
			wg.Done()
		}(mx, val)
	}

	wg.Wait()
	fmt.Fprintln(os.Stdout)
}

func CountArrayChannel() {
	res := make(chan int, len(numbers))

	go func(ch chan<- int) {
		for _, val := range numbers {
			ch <- val * val
		}
		// закрываем канал по окончанию записи
		close(ch)
	}(res)

	for {
		data, ok := <-res
		// Читаем из канала пока он открыт
		if ok {
			fmt.Println(data)
		} else {
			return
		}
	}
}

func CountArrayChannelWg() {
	// Задаем размер массива чисел
	size := len(numbers)

	wg := &sync.WaitGroup{}
	// Создаем канал с буфером == size
	res := make(chan int, size)

	wg.Add(size)

	for _, val := range numbers {
		go func(ch chan<- int, n int) {
			ch <- n * n
			wg.Done()
		}(res, val)
	}

	// По завершению записи закрываем канал
	wg.Wait()
	close(res)

	// Читаем из канала пока тот открыт
	for data := range res {
		fmt.Println(data)
	}
}