package tree

import (
	"math/big"
)

// algorithm "long numbers of approximately the same length multiply more efficiently"
// from https://habr.com/ru/post/255761/
// details https://hsto.org/files/292/9a7/c6f/2929a7c6f0e04cccb1e30b9417b1492b.png

func factorialTree(l uint64, r uint64) *big.Int {
	if l > r {
		return big.NewInt(1)
	}

	if l == r {
		return new(big.Int).SetUint64(l)
	}

	if r-l == 1 {
		return new(big.Int).SetUint64(l * r)
	}
	// calculating the subtree
	m := (l + r) / 2
	// TODO (yusupov) add goroutine
	result := factorialTree(l, m)
	result.Mul(result, factorialTree(m+1, r))
	return result
}

func FactorialTree(n uint64) *big.Int {
	if n == 0 {
		return big.NewInt(1)
	} else if n == 1 || n == 2 {
		return new(big.Int).SetUint64(n)
	}
	return factorialTree(2, n)
}
