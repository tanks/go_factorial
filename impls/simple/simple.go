package simple

import (
	"math/big"
)

func FactorialSimple(n uint64) (result *big.Int) {
	result = new(big.Int)
	result.SetInt64(1)
	if n > 0 {
		result.SetUint64(n)
		result.Mul(result, FactorialSimple(n-1))
	}
	return
}
