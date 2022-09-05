package main

import "fmt"

func removeItem(numbers []int, element int) []int {
	return append(numbers[:element], numbers[element+1:]...)
}

func main() {
	slice := []int{1, 5, 10, 22, 31, 45, 12, 11, 45}
	fmt.Println("Before:", slice)
	fmt.Println("After: ", removeItem(slice, 3))
}