package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var currentFunc = func(s string) int {
	return len(s)
}

var improvedFunc = func(s string) int {
	return strings.Count(s, "") - 1
}

var scaling = 50

func runCharCountExperiment(s string) int {
	return newExperiment(currentFunc, improvedFunc, 100 * time.Millisecond).run(s)
}

func newExperiment(current, improved func(string) int, timeout time.Duration) *experiment {
	return &experiment{
		current,
		improved,
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
		select {
		case c <- &funcResult{
			ex.improved(s),
			time.Since(st),
		}:
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

	// use metrics iso prints
	fmt.Printf("Current functionality duration: %s\n", curDuration)
	fmt.Printf("Improved functionality duration: %s\n", imprDuration)

	if cur != impr {
		fmt.Printf("ERROR current result != improvement result: %+v != %+v\n", cur, impr)
		return cur
	}

	// scale improved func up (or down)
	if e.r.Intn(100) < scaling {
		fmt.Printf("Chose improved functionality\n")
		return impr
	}

	fmt.Printf("Chose current functionality\n")
	return cur
}