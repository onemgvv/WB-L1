package main

import (
	"fmt"
	"strings"
)

func reverse1(word string) string {
	// разбиваем строку на массив рун
	runes := []rune(word)

	/*
		проходим по массиву с двух сторон
		i - с начала
		j - с конца
		и меняем местами символы
	*/
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// преобразуем руны в string и возвращаем
	return string(runes)
}

func reverse2(word string) string {
	// разбиваем строку на массив 
	arr := strings.Split(word, "")
	var reverse []string
	
	// разворачиваем массив в цикле
	for i := len(arr) - 1; i >= 0; i-- {
		reverse = append(reverse, arr[i])
	}

	// собираем строку из массива и возвращаем
	return strings.Join(reverse, "")
}

func main() {
	fmt.Println("Reverse as rune: ", reverse1("главрыба"))
	fmt.Println("Reverse as array split join: ", reverse2("главрыба"))
}