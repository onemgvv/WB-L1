package worker

import (
	"fmt"
	"sync"
)

func StartWorker(wg *sync.WaitGroup, id int, c chan int) {
	// Выводим данные пока канал не закрыт
	for res := range c {
		fmt.Println("ID:", id, "RES:", res)
	}

	// сообщаем wg об остановке горутины
	wg.Done() 
}