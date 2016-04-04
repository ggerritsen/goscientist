package goscientist

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"
)

type ExperimentError error

func NewExperiment(currentFunc, improvedFunc interface{}) (*experiment, ExperimentError) {
	c := reflect.ValueOf(currentFunc)
	if c.Kind() != reflect.Func {
		return nil, fmt.Errorf("currentFunc is not a function")
	}
	i := reflect.ValueOf(improvedFunc)
	if i.Kind() != reflect.Func {
		return nil, fmt.Errorf("improvedFunc is not a function")
	}

	return &experiment{
		c,
		i,
		100 * time.Millisecond,
		50, // choose improved functionality half of the time
		rand.New(rand.NewSource(time.Now().UnixNano())),
	}, nil
}

type experiment struct {
	currentFunc  reflect.Value
	improvedFunc reflect.Value
	timeout      time.Duration
	scaling      int
	r            *rand.Rand
}

type funcResult struct {
	result   []interface{}
	duration time.Duration
}

// run runs the experiment and returns the return values of the
// functions that are in the experiment, in combination with an
// error (or nil).
func (e *experiment) Run(s ...interface{}) ([]interface{}, ExperimentError) {
	// check input first
	if len(s) != e.currentFunc.Type().NumIn() {
		return nil, fmt.Errorf("Number of inputs (%d) is incorrect", len(s))
	}

	input := make([]reflect.Value, len(s))
	for i, arg := range s {
		if t1, t2 := reflect.ValueOf(arg).Type(), e.currentFunc.Type().In(i); t1 != t2 {
			return nil, fmt.Errorf("Incorrect argument type (%d): %s. Expected type %s", i, t1, t2)
		}
		input[i] = reflect.ValueOf(arg)
	}

	c := make(chan *funcResult, 1)
	go func(ex *experiment) {
		start := time.Now()
		r := ex.improvedFunc.Call(input)
		d := time.Since(start)

		output := make([]interface{}, len(r))
		for i, a := range r {
			output[i] = a.Interface()
		}
		res := &funcResult{
			output,
			d,
		}

		c <- res
	}(e)

	start := time.Now()
	r := e.currentFunc.Call(input)
	curDuration := time.Since(start)

	// TODO perhaps use reflect.ValueOf(e.currentFunc).Type().Out()?
	output := make([]interface{}, len(r))
	for i, a := range r {
		output[i] = a.Interface()
	}

	var funcRes *funcResult
	select {
	case funcRes = <-c:
	case <-time.After(e.timeout):
	}

	if funcRes == nil {
		log.Printf("No experiment outcome, improved func took too long\n")
		return output, nil
	}

	impr := funcRes.result
	imprDuration := funcRes.duration

	// TODO use metrics iso printouts
	if imprDuration < curDuration {
		log.Printf("Improved faster (%s vs. %s)", curDuration, imprDuration)
	} else {
		log.Printf("Current faster (%s vs. %s)", curDuration, imprDuration)
	}
	//	fmt.Printf("Current functionality duration: %s\n", curDuration)
	//	fmt.Printf("Improved functionality duration: %s\n", imprDuration)

	if !eq(output, impr) {
		log.Printf("ERROR current result != improvement result for input %+v, choosing current functionality (%+v != %+v)\n", s, output, impr)
		return output, nil
	}

	// TODO scale improved func up (or down)
	if e.r.Intn(100) < e.scaling {
		log.Printf("Chose improved functionality\n")
		return impr, nil
	}

	log.Printf("Chose current functionality\n")
	return output, nil
}

func eq(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i, aa := range a {
		if !reflect.DeepEqual(aa, b[i]) {
			return false
		}
	}
	return true
}
