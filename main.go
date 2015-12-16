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
	//	timeout := 2 * time.Second

	var legacyResult, shinyNewResult int
	var legacyDuration, shinyNewDuration time.Duration
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		legacyResult = legacyFunc(s)
		legacyDuration = time.Since(start)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		shinyNewResult = shinyNewFunc(s)
		shinyNewDuration = time.Since(start)
		wg.Done()
	}()
	wg.Wait() // add timeout through context?

	if legacyResult != shinyNewResult {
		fmt.Printf("ERROR legacy result %+v != shiny new result %+v\n", legacyResult, shinyNewResult)
	}

	fmt.Printf("Legacy duration: %s\n", legacyDuration)
	fmt.Printf("Shiny new duration: %s\n", shinyNewDuration)

	return legacyResult
}
