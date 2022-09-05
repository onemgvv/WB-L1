package main

import (
	"fmt"
	"l1/03/sqrt_sum"
)

func main() {
	fmt.Println("Sum Numbers Square")
	fmt.Println()
	fmt.Println("Sum Numbers Square (Channel):")
	sqrtsum.SqrtSumChan()
	fmt.Println("Sum Numbers Square (Mutex):")
	sqrtsum.SqrtSumMutex()
	fmt.Println("Sum Numbers Square (Atomic):")
	sqrtsum.SqrtSumAtomic()
}