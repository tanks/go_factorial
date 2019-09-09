package main

import (
	"awesomeProject/impls/simple/impls/simple"
	"awesomeProject/impls/simple/impls/tree"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	impls := []struct {
		name string
		fun  FactType
	}{
		{"simple", simple.FactorialSimple},
		{"tree", tree.FactorialTree},
	}
	for _, funct := range impls {
		b.Run(funct.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				funct.fun(2000)
			}
		})
	}
}
