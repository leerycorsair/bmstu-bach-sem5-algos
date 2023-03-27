package ants

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func GenMap(cities int) MatrixInt {
	currMap := mapInit(cities, cities)

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < cities; i++ {
		for j := i + 1; j < cities; j++ {
			currMap[i][j] = rand.Intn(50) + 1
			currMap[j][i] = currMap[i][j]
		}
	}
	return currMap
}

func PrintMap(citiesMap MatrixInt) {
	for i := 0; i < len(citiesMap); i++ {
		for j := 0; j < len(citiesMap); j++ {
			fmt.Printf("%5d", citiesMap[i][j])
		}
		fmt.Println()
	}
}

func LatexMap(citiesMap MatrixInt, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; i < len(citiesMap); i++ {
		for j := 0; j < len(citiesMap)-1; j++ {
			fmt.Fprintf(file, "%5d &", citiesMap[i][j])
		}
		fmt.Fprintf(file, "%5d \\\\ \n", citiesMap[i][len(citiesMap)-1])
	}
}

func mapInit(rows int, colunms int) MatrixInt {
	var mtr = make([][]int, rows)
	for i := range mtr {
		mtr[i] = make([]int, colunms)
	}
	return mtr
}
