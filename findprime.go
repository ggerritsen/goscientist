package main

import (
	"log"
	"math/big"
)

var primeCurrent = func(n int) bool {
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

var primeImproved = func(n int) bool {
	i := big.NewInt(int64(n))
	return i.ProbablyPrime(1)
}

func runFindPrimeExperiment(n int) bool {
	exp, err := newExperiment(primeCurrent, primeImproved)
	if err != nil {
		log.Printf("Error: could not create new experiment: %s", err)
		return primeCurrent(n)
	}

	result, err := exp.run(n)
	if err != nil {
		log.Printf("Could not run experiment: %s", err)
		return primeCurrent(n)
	}
	return result[0].(bool)
}
