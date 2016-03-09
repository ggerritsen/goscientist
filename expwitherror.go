package main

import (
	"log"
	"strconv"
)

func parseIntCurrent(s string) (int, error) {
	return strconv.Atoi(s)
}

func parseIntImproved(s string) (int, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(i), nil
}

func runParseIntExperiment(s string) (int, error) {
	exp, err := newExperiment(parseIntCurrent, parseIntImproved)
	if err != nil {
		log.Printf("Error: could not create new experiment: %s", err)
		return parseIntCurrent(s)
	}

	r, err := exp.run(s)
	if err != nil {
		log.Printf("Could not run experiment: %s", err)
		return parseIntCurrent(s)
	}

	i := r[0].(int)
	var e error
	if x := r[1]; x != nil {
		e = x.(error)
	}
	return i, e
}
