package main

import (
	"fmt"
	"sync"
	"time"
)

var legacyFunc = func(s string) int {
	println("legacy: " + s)
	return 2
}

var shinyNewFunc = func(s string) int {
	println("shinyNew: " + s)
	return 3
}

func main() {
	input := []string{"Ajax", "PSV", "Feyenoord"}

	for _, s := range input {
		expFunc(s)
	}

	println("done.")
}

func expFunc(s string) int {
	exp := &experiment{}
	exp.current = legacyFunc
	exp.improvement = shinyNewFunc

	return exp.run(s)
}

type experiment struct {
	current func(s string) int
	improvement func(s string) int
	// timeout
}

func (e *experiment) run(s string) int {
	// add scaling

	var cur, impr int
	var curDuration, imprDuration time.Duration
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		cur = e.current(s)
		curDuration = time.Since(start)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		impr = shinyNewFunc(s)
		imprDuration = time.Since(start)
	}()
	wg.Wait() // add timeout through context?

	if cur != impr {
		fmt.Printf("ERROR current result != improvement result: %+v != %+v\n", cur, impr)
	}

	// use metrics iso prints
	fmt.Printf("Current functionality duration: %s\n", curDuration)
	fmt.Printf("Improved functionality duration: %s\n", imprDuration)

	return cur
}