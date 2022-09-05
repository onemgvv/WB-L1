package main

import (
	"fmt"
	"strings"
)

func flipper(sentence string) string {
	// разбиваем предложение на массив строк по пробелу
	arr := strings.Split(sentence, " ")
	var rev []string

	// разворачиваем массив в цикле
	for i := len(arr) - 1; i >= 0; i-- {
		rev = append(rev, arr[i])
	}

	// "склеиваем" предложение из массива
	return strings.Join(rev, " ")
}

func main() {
	fmt.Println(flipper("snow dog sun"))
}