package main

import "fmt"

var (
	count1 int
	count2 int
)

func binarySearch(array []int, item int) int {
	var start int 	  // Стартовая позиция 0
	end := len(array) // Конечная len(array)
	var middle int
	found := false
	position := -1		// если элемен не найден возвращаем -1

	// Пока элемент не найден и стартовая позиция меньше конечной
	// Вычисляем середину массива
	// если центральный эемент равен искомому устанавливаем found в true
	// и возвращаем позицию найденого элемента
	for found == false && start <= end {
		count1++
		middle = (start + end) / 2
		if array[middle] == item {
			found = true
			position = middle
			return position
		}

		// Если искомый элемент меньше центрального меняем конечную позицию
		// В противном случае меняем начальную
		if item < array[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return position
}

func recursiveBinarySearch(array []int, item, start, end int) int {
	// Если стартовая позиция больше конечно возвращаем -1
	if start > end {
		return -1
	}

	// Находим середину массива
	middle := (start + end) / 2
	
	count2++

	// Если искомый элемент посередине возвращаем его позицию
	if item == array[middle] {
		return middle
	}

	// Если искомый элемент меньше среднего, вызываем функцию повторно изменив конечную позицию
	// В противном случае меняем стартовую позицию и вызываем функцию повторно
	if item < array[middle] {
		return recursiveBinarySearch(array, item, start, middle-1)
	} else {
		return recursiveBinarySearch(array, item, middle+1, end)
	}
}

func main() {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Binary search in loop")
	fmt.Println(binarySearch(array, 9))
	fmt.Println(count1) // Количество итераций в цикле
	fmt.Println("Binary search in recursive fn")
	fmt.Println(recursiveBinarySearch(array, 9, 0, len(array)))
	fmt.Println(count2) // Количество итераций в рекурсии
}