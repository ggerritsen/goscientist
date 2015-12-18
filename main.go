package main
import "fmt"

func main() {
	clubs := []string{"Ajax", "PSV", "Feyenoord"}
	numbers := []int{1,2,3,4}

	for _, s := range clubs {
		runCharCountExperiment(s)
	}
	fmt.Printf("Char count experiment done.\n\n")

	for _, s := range clubs {
		runFindChampionExperiment(s)
	}
	fmt.Printf("Find champion experiment done.\n\n")

	for _, n := range numbers {
		runFindPrimeExperiment(n)
	}
	fmt.Printf("Find prime experiment done.\n")
}