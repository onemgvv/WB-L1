package main

import (
	"bignum"
	"fmt"
)

func main() {
	bNumber := bignum.NewBigNumbers(724389233415407581, 292301003375467513)
	fmt.Printf("Add: %d\nSubstract: %d\nMultiple: %d\nDivide: %d\n", bNumber.Add(), bNumber.Substact(), bNumber.Multiple(), bNumber.Divide())
}