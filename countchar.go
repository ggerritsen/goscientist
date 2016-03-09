package main

import (
	"log"
	"strings"
)

var charCurrent = func(s string) int {
	return strings.Count(s, "") - 1
}

var charImproved = func(s string) int {
	return len(s)
}

func runCharCountExperiment(s string) int {
	exp, err := newExperiment(charCurrent, charImproved)
	if err != nil {
		log.Printf("Error: could not create new experiment: %s", err)
		return charCurrent(s)
	}

	result, err := exp.run(s)
	if err != nil {
		log.Printf("Could not run experiment: %s", err)
		return charCurrent(s)
	}
	return result[0].(int)
}