package main

import (
	"fmt"
	"strings"
)

var justString string
func someFunc() string {
  v := createHugeString(1 << 10)
  justString = v[:100]
  return justString
}

func createHugeString(l int) string {
  return strings.Repeat("a", l)
}

func main() {
  res := someFunc()
  fmt.Println(res)
}