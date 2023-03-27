package main

import (
	"fmt"
	"local/levenshtein"
)

func main() {
	var f_str, s_str string
	fmt.Println("Enter the first string:")
	fmt.Scanln(&f_str)
	fmt.Println("Enter the second string:")
	fmt.Scanln(&s_str)

	dist := levenshtein.LevRec(f_str, s_str)
	fmt.Printf("Recursive Levenshtein: %d\n", dist)

	// dist, mtr := levenshtein.LevMtr(f_str, s_str)
	// fmt.Printf("Matrix Levenshtein: %d\n", dist)
	// levenshtein.Print_matrix(mtr)

	dist = levenshtein.LevRecCached(f_str, s_str)
	fmt.Printf("Cached Levenshtein: %d\n", dist)

	dist = levenshtein.DamLevRec(f_str, s_str)
	fmt.Printf("Recursive Damerau-Levenshtein: %d\n", dist)

	dist, mtr := levenshtein.DamLevMtr(f_str, s_str)
	fmt.Printf("Matrix Damerau-Levenshtein: %d\n", dist)
	levenshtein.Print_matrix(mtr)

}
