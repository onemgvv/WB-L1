package main

import (
	"fmt"
	"l1/12/set"
)

func main() {
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	newSet := set.Set{}

	for _, v := range data {
		newSet.Add(v)
	}

	fmt.Println(newSet)
}