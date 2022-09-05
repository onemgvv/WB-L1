package main

import (
	"fmt"
	"l1/24/pointer"
)

func main() {
	point := pointer.NewPointer(13.3, 26)
	fmt.Println(point.CalcDistance())
}