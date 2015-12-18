package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var legacyFunc = func(s string) int {
	return 2
}

var shinyNewFunc = func(s string) int {
	return 99999
}

var scaling = 0

func main() {
	scaling = 10
	input := []string{"Ajax", "PSV", "Feyenoord"}

	for _, s := range input {
		runExperiment(s)
	}

	println("done.")
}

func runExperiment(s string) int {
	return newExperiment(legacyFunc, shinyNewFunc).run(s)
}

func newExperiment(current, improvement func(string) int) *experiment {
	return &experiment{
		current,
		improvement,
		rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type experiment struct {
	current func(s string) int
	improvement func(s string) int
	r *rand.Rand
	// timeout
}

func (e *experiment) run(s string) int {
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

	// scale improved func up (or down)
	if e.r.Intn(100) < scaling {
		return impr
	}

	return cur
}