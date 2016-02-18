package main
import "fmt"

func main() {
	clubs := []string{"Ajax", "PSV", "Feyenoord"}
	numbers := []int{1,2,3,4}

	fmt.Printf("Starting char count experiment.\n")
	for _, s := range clubs {
		runCharCountExperiment(s)
	}
	fmt.Printf("Char count experiment done.\n\n")

	fmt.Printf("Starting find champion experiment.\n")
	for _, s := range clubs {
		runFindChampionExperiment(s)
	}
	fmt.Printf("Find champion experiment done.\n\n")

	fmt.Printf("Starting find prime experiment.\n")
	for _, n := range numbers {
		runFindPrimeExperiment(n)
	}
	fmt.Printf("Find prime experiment done.\n")
}