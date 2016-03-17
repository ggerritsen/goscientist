package main

import (
	"log"
	"strings"

	"github.com/ggerritsen/goscientist"
)

var championCurrent = func(s string) bool {
	return s[0] == 'A'
}

var championImproved = func(s string) bool {
	return strings.ToLower(s) == "ajax"
}

func runFindChampionExperiment(s string) bool {
	exp, err := goscientist.NewExperiment(championCurrent, championImproved)
	if err != nil {
		log.Printf("Error: could not create new experiment: %s", err)
		return championCurrent(s)
	}

	result, err := exp.Run(s)
	if err != nil {
		log.Printf("Could not run experiment: %s", err)
		return championCurrent(s)
	}
	return result[0].(bool)
}
