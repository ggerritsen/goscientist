package main

import (
	"strings"
)

var charCurrentFunc = func(s interface{}) interface{} {
	input := s.(string)
	return len(input)
}

var charImprovedFunc = func(s interface{}) interface{} {
	input := s.(string)
	return strings.Count(input, "") - 1
}

func runCharCountExperiment(s string) int {
	result := newExperiment(charCurrentFunc, charImprovedFunc).run(s)
	return result.(int)
}