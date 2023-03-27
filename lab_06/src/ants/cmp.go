package ants

import (
	"fmt"
	"log"
	"os"
)

const (
	CITIES = 10

	MIN_ALPHA = 1.0
	MAX_ALPHA = 3.0

	MIN_P = 0.2
	MAX_P = 1
	STEP  = 0.2

	MIN_DAYS = 9
	MAX_DAYS = 11
)

func CompareAlgos() {
	citiesMap := GenMap(CITIES)
	LatexMap(citiesMap, "./mtrFile.txt")

	cmpFile, err := os.Create("./cmpFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer cmpFile.Close()

	dist_B, _ := BruteForceAlgo(citiesMap, CITIES)
	for alpha := MIN_ALPHA; alpha <= MAX_ALPHA; alpha++ {
		for p := MIN_P; p <= MAX_P; p += STEP {
			for days := MIN_DAYS; days <= MAX_DAYS; days++ {
				dist_A, _ := AntAlgo(citiesMap, alpha, DEFAULT_BETA, p, DEFAULT_PHEROMONCOEF, days)
				fmt.Fprintf(cmpFile, "%5.2v & %5.2v & %5v & %5d \\\\ \\hline \n", alpha, p, days, absDiffInt(dist_A, dist_B))
			}
		}
	}

}
