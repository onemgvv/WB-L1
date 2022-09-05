package main

import "fmt"

func main() {
	var num1, num2 int = 5, 10
	fmt.Printf("Num1: %d; Num2: %d\n", num1, num2)
	num1, num2 = num2, num1
	fmt.Printf("Num1: %d; Num2: %d\n", num1, num2)
}