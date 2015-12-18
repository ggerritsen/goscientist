package main

import (
	"strings"
)

var championCurrentFunc = func(s interface{}) interface{} {
	input := s.(string)
	return strings.ToLower(input) == "ajax"
}

var championImprovedFunc = func(s interface{}) interface{} {
	input := s.(string)
	return input[0] == 'A'
}

func runFindChampionExperiment(s string) bool {
	result := newExperiment(championCurrentFunc, championImprovedFunc).run(s)
	return result.(bool)
}
