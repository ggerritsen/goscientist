package main

import (
	"math/big"
	"log"
)

var primeCurrentFunc = func(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

var primeImprovedFunc = func(n int) bool {
	i := big.NewInt(int64(n))
	return i.ProbablyPrime(1)
}

func runFindPrimeExperiment(n int) bool {
	exp, err := newExperiment(primeCurrentFunc, primeImprovedFunc)
	if err != nil {
		log.Fatalf("Could not create new experiment: %s", err)
	}

	result, err := exp.run(n)
	if err != nil {
		log.Fatalf("Could not run experiment: %s", err)
	}
	return result[0].(bool)
}
