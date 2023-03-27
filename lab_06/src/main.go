package main

import (
	"fmt"
	"local/ants"
)

func main() {
	fmt.Println("Ant Algorithm VS Brute Force\n")
	var cities int
	fmt.Printf("Cities:")
	fmt.Scanln(&cities)

	citiesMap := ants.GenMap(cities)

	fmt.Printf("The Adjacency Matrix:\n")
	ants.PrintMap(citiesMap)

	dist, path := ants.AntAlgo(citiesMap, ants.DEFAULT_ALPHA, ants.DEFAULT_BETA,
		ants.DEFAULT_P, ants.DEFAULT_PHEROMONCOEF, ants.DEFAULT_DAYS)
	fmt.Printf("%25s: LEN = %5d, PATH = %v\n", "Ant Algorithm", dist, path)

	dist, path = ants.BruteForceAlgo(citiesMap, cities)
	fmt.Printf("%25s: LEN = %5d, PATH = %v\n", "Brute Force Algorithm", dist, path)

	ants.CompareAlgos()
}
