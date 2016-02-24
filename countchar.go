package main

import (
	"strings"
	"log"
)

var charCurrentFunc = func(s string) int {
	return len(s)
}

var charImprovedFunc = func(s string) int {
	return strings.Count(s, "") - 1
}

func runCharCountExperiment(s string) int {
	exp, err := newExperiment(charCurrentFunc, charImprovedFunc)
	if err != nil {
		log.Fatalf("Could not create new experiment: %s", err)
	}

	result, err := exp.run(s)
	if err != nil {
		log.Fatalf("Could not run experiment: %s", err)
	}
	return result[0].(int)
}

// TODO perhaps use interface?