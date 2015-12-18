package main

import (
	"fmt"
	"math/rand"
	"time"
)

var legacyFunc = func(s string) int {
	return 2
}

var shinyNewFunc = func(s string) int {
	time.Sleep(10*time.Millisecond)
	return 2
}

var scaling = 50

func main() {
	input := []string{"Ajax", "PSV", "Feyenoord"}

	for _, s := range input {
		runExperiment(s)
	}

	println("done.")
}

func runExperiment(s string) int {
	return newExperiment(legacyFunc, shinyNewFunc, 100 * time.Millisecond).run(s)
}

func newExperiment(current, improvement func(string) int, timeout time.Duration) *experiment {
	return &experiment{
		current,
		improvement,
		timeout,
		rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type experiment struct {
	current func(s string) int
	improved func(s string) int
	timeout time.Duration
	r *rand.Rand
}

type funcResult struct {
	result int
	duration time.Duration
}

func (e *experiment) run(s string) int {
	var cur int
	var curDuration time.Duration
	start := time.Now()

	c := make(chan *funcResult, 1)
	go func(ex *experiment, st time.Time) {
		r := ex.improved(s)
		dur := time.Since(st)
		select {
			case c <- &funcResult{r, dur}:
			default:
		}
	}(e, start)

	cur = e.current(s)
	curDuration = time.Since(start)

	var funcRes *funcResult
	select {
	case funcRes = <- c:
	case <- time.After(e.timeout):
	}

	if funcRes == nil {
		fmt.Printf("No experiment outcome, improved func took too long\n")
		return cur
	}

	impr := funcRes.result
	imprDuration := funcRes.duration

	if cur != impr {
		fmt.Printf("ERROR current result != improvement result: %+v != %+v\n", cur, impr)
	}

	// use metrics iso prints
	fmt.Printf("Current functionality duration: %s\n", curDuration)
	fmt.Printf("Improved functionality duration: %s\n", imprDuration)

	// scale improved func up (or down)
	if e.r.Intn(100) < scaling {
		fmt.Printf("Chose improved functionality\n")
		return impr
	}

	fmt.Printf("Chose current functionality\n")
	return cur
}