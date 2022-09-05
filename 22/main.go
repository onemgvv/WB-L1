package main

import (
	"fmt"
	"math/big"
)

type BigNumberWorker interface {
	Add() *big.Int
	Substact() *big.Int
	Multiple() *big.Int
	Divide() *big.Int
}

type BigNumbers struct {
	n1, n2 *big.Int
}

func (b *BigNumbers) Add() *big.Int {
	return new(big.Int).Add(b.n1, b.n2)
}

func (b *BigNumbers) Substact() *big.Int {
	return new(big.Int).Sub(b.n1, b.n2)
}

func (b *BigNumbers) Multiple() *big.Int {
	return new(big.Int).Mul(b.n1, b.n2)
}

func (b *BigNumbers) Divide() *big.Int {
	return new(big.Int).Div(b.n1, b.n2)
}

func NewBigNumbers(n1, n2 int64) BigNumberWorker {
	return &BigNumbers{big.NewInt(n1), big.NewInt(n2)}	
}

func main() {
	bNumber := NewBigNumbers(724389233415407581, 292301003375467513)
	fmt.Printf("Add: %d\nSubstract: %d\nMultiple: %d\nDivide: %d\n", bNumber.Add(), bNumber.Substact(), bNumber.Multiple(), bNumber.Divide())
}