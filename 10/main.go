package main

import (
	"fmt"
	"math"
)

func main() {
	temps := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temp := make(map[int][]float64)

	for _, i := range temps {
		if(i > 0) {
			step := int(math.Floor(i / 10))
			temp[step * 10] = append(temp[step * 10], i)
		} else {
			step := int(math.Ceil(i / 10))
			temp[step * 10] = append(temp[step * 10], i)
		}
	}

	fmt.Println(temp)
}