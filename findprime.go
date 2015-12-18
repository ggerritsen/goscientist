package main

import (
	"math/big"
)

var primeCurrentFunc = func(n interface{}) interface{} {
	input := n.(int)

	if input <= 1 {
		return false
	}
	for i := 2; i < input; i++ {
		if input % i == 0 {
			return false
		}
	}

	return true
}

var primeImprovedFunc = func(n interface{}) interface{} {
	input := n.(int)

	i := big.NewInt(int64(input))
	return i.ProbablyPrime(1)
}

func runFindPrimeExperiment(n int) bool {
	result := newExperiment(primeCurrentFunc, primeImprovedFunc).run(n)
	return result.(bool)
}
