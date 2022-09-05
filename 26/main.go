package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type LetterSet map[rune]uint

func checkUnique(sentence string) bool {
	ltSet := LetterSet{}
	lwSent := strings.ToLower(sentence)

	// преобразуем строку в массив рун
	// и проэтерируемся по нему
	for _, v := range []rune(lwSent) {
		// добавляем увеличиваем количество каждого символа в мапе
		ltSet[v] += 1

		// если хоть один символ в мапе больше 1
		// возвращаем false
		if(ltSet[v] > 1) {
			return false
		}
	}

	// если все символы в мапе в количества 1шт возвращаем true
	return true
}

func main() {
	var word string

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	
	defer out.Flush()

	fmt.Println("Введите слово для проверки:")

	fmt.Fscanf(in, "%s", &word)
	fmt.Fprint(out, "unique:", checkUnique(word))
}