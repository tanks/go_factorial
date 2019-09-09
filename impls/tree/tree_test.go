package tree

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	t.Run("n = 0", func(t *testing.T) {
		got := FactorialTree(0)
		var wantRes *big.Int = big.NewInt(1)
		if got.Cmp(wantRes) != 0 {
			t.Errorf("factorial(1) = %d; want %d", got, wantRes)
		}
	})
	t.Run("n = 1", func(t *testing.T) {
		got := FactorialTree(1)
		var wantRes *big.Int = big.NewInt(1)
		if got.Cmp(wantRes) != 0 {
			t.Errorf("factorial(1) = %d; want %d", got, wantRes)
		}
	})

	t.Run("n = 3", func(t *testing.T) {
		got := FactorialTree(3)
		var wantRes *big.Int = big.NewInt(6)
		if got.Cmp(wantRes) != 0 {
			t.Errorf("factorial(1) = %d; want %d", got, wantRes)
		}
	})

	t.Run("n = 4", func(t *testing.T) {
		got := FactorialTree(4)
		var wantRes *big.Int = big.NewInt(24)
		if got.Cmp(wantRes) != 0 {
			t.Errorf("factorial(1) = %d; want %d", got, wantRes)
		}
	})

	t.Run("n = 16", func(t *testing.T) {
		got := FactorialTree(16)
		var wantRes *big.Int = big.NewInt(20922789888000)
		if got.Cmp(wantRes) != 0 {
			t.Errorf("factorial(1) = %d; want %d", got, wantRes)
		}
	})
	t.Run("n = 20", func(t *testing.T) {
		got := FactorialTree(20)
		var wantRes *big.Int = big.NewInt(2432902008176640000)
		if got.Cmp(wantRes) != 0 {
			t.Errorf("factorial(1) = %d; want %d", got, wantRes)
		}
	})
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// perform the operation we're analyzing
		FactorialTree(2000)
	}
}
