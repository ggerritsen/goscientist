package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newExperiment(current, improved func(interface{}) interface{}) *experiment {
	return &experiment{
		current,
		improved,
		100 * time.Millisecond,
		50, // choose improved functionality half of the time
		rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type experiment struct {
	current func(interface{}) interface{}
	improved func(interface{}) interface{}
	timeout time.Duration
	scaling int
	r *rand.Rand
}

type funcResult struct {
	result interface{}
	duration time.Duration
}

func (e *experiment) run(s interface{}) interface{} {
	var cur interface{}
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

	// use metrics iso printouts
	fmt.Printf("Current functionality duration: %s\n", curDuration)
	fmt.Printf("Improved functionality duration: %s\n", imprDuration)

	if cur != impr {
		fmt.Printf("ERROR current result != improvement result for input %+v, choosing current functionality (%+v != %+v)\n", s, cur, impr)
		return cur
	}

	// scale improved func up (or down)
	if e.r.Intn(100) < e.scaling {
		fmt.Printf("Chose improved functionality\n")
		return impr
	}

	fmt.Printf("Chose current functionality\n")
	return cur
}