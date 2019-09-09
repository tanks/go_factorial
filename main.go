package main

import (
	"awesomeProject/impls/simple/impls/simple"
	"awesomeProject/impls/simple/impls/tree"
	"fmt"
	"math/big"
)

type FactType func(n uint64) *big.Int

func factorial(n uint64, impl FactType) *big.Int {
	return impl(n)
}

func main() {
	var n uint64
	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&n)
	fmt.Println("Factorial calc simple is: \t\t", factorial(n, simple.FactorialSimple))
	fmt.Println("Factorial calc tree is: \t\t", factorial(n, tree.FactorialTree))

}
