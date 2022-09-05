package main

import "fmt"

var count int

func quickSort(array []int) []int {
	result := []int{}

	// Если длина массива меньше 1 возвращаем его
	// Базовый случай для выхода из рекурсии
	if len(array) < 1 {
		return array
	}

	// Получаем индекс опорного жлемента
	pivotIndex := len(array) / 2
	// Получаем сам элемент по индексу
	pivot := array[pivotIndex]
	// Массив меньших элементов
	var less []int
	// Массив больших элементов
	var greater []int

	// Проходимся в цикле по массиву
	for i := 0; i < len(array); i++ {
		count++
		// Если индекс равен опорному пропускаем шаг
		if i == pivotIndex {
			continue
		}

		// Если элемент массива меньше опорного
		// добавляем его в массив меньших элементов
		// Если больше - в массив больших
		if array[i] < pivot {
			less = append(less, array[i])
		} else {
			greater = append(greater, array[i])
		}
	}

	// Разворачиваем в переменную result результат функции quickSort от массива меньших значений
	result = append(result, quickSort(less)...)
	// Затем добавляем опорный элемент
	result = append(result, pivot)
	// И наконец разворачиваем результат функции от массива больших значений
	result = append(result, quickSort(greater)...)
	return result
}

func main() {
	array := []int{9, 10, 0, 5, 19, 8, 7, 15, 12, 16, 23, 25, 27, 30, 1, 3, 6, 4}
	fmt.Println(quickSort(array))
	fmt.Println(count) // Количество итераций
}