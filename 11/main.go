package main

import (
	"fmt"
)

func getIntersection(a []string, b []string, mode byte) []string {
	m := make(map[string]byte)

	fmt.Println(m)

	for _, k := range a {
		m[k] += 1
	}

	for _, k := range b {
		m[k] += 2
	}

	fmt.Println(m)

	result := []string{}

	for k, v := range m {
		if v == mode {
			result = append(result, k)
		}
	}

	return result
}

func main() {

	a := []string{"go", "java", "rust", "nodejs"}
	b := []string{"c++", "go", "nodejs", "c#"}

	fmt.Println(getIntersection(a, b, 3))
}
