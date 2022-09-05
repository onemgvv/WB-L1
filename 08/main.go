package main

import (
	"errors"
	"fmt"
	"io"
	"log"
)

func ReadConsole(n *int8) {
	_, ok := fmt.Scan(n)
	if ok != nil {
		if ok != io.EOF {
			log.Fatalln(ok)
		} else {
			fmt.Println("exit")
		}
		return
	}
}

func SwapByte(data *int64, n int8) error {
	if n > 64 || n < 1 {
		return errors.New("Русским языком написал, от 1 до 64!")
	}
	*data ^= 1 << (n - 1)
	return nil
}

func PrintBytes(data int64) {
	fmt.Print("Значение: ", data, "\nБайты: ")
	for i := 0; i < 64; i++ {
		if (data & (1 << i)) == 0 {
			print("0")
		} else {
			print("1")
		}
	}
	print("\n")
}

func main() {
	var n int8
	var data int64
	fmt.Println("Введите номер байта [1 - 64]: ")
	for {
		//Получаем значение с консоли
		ReadConsole(&n)

		// Меняем бит в соответствии номером
		ok := SwapByte(&data, n)
		if ok != nil {
			log.Println(ok, "Ошибка, повтори попытку")
		}

		// Выводим результат
		PrintBytes(data)
	}
}