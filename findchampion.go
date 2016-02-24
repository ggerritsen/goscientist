package main

import (
	"strings"
	"log"
)

var championCurrentFunc = func(s string) bool {
	return strings.ToLower(s) == "ajax"
}

var championImprovedFunc = func(s string) bool {
	return s[0] == 'A'
}

func runFindChampionExperiment(s string) bool {
	exp, err := newExperiment(championCurrentFunc, championImprovedFunc)
	if err != nil {
		log.Fatalf("Could not create new experiment: %s", err)
	}

	result, err := exp.run(s)
	if err != nil {
		log.Fatalf("Could not run experiment: %s", err)
	}
	return result[0].(bool)
}
