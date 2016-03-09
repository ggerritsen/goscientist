package main

import (
	"log"
)

func main() {
	clubs := []string{"Ajax", "PSV", "Feyenoord", "Groningen", "Telstar"}
	numbers := []int{1, 2, 3, 4}
	str := []string{"1", "2", "a", "%"}

	log.Println("Starting char count experiment.")
	for _, s := range clubs {
		runCharCountExperiment(s)
	}
	log.Println("Char count experiment done.\n")

	log.Println("Starting find champion experiment.")
	for _, s := range clubs {
		runFindChampionExperiment(s)
	}

	log.Println("Starting find prime experiment")
	for _, n := range numbers {
		runFindPrimeExperiment(n)
	}
	log.Println("Find prime experiment done\n")

	log.Println("Starting parse int experiment")
	for _, s := range str {
		runParseIntExperiment(s)
	}
	log.Println("Parse int experiment done.")
}