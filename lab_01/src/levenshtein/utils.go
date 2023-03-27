package levenshtein

import "fmt"

func min_value(args ...int) int {
	var min int = args[0]
	for _, value := range args {
		if min > value {
			min = value
		}
	}
	return min
}

func max_value(args ...int) int {
	var max int = args[0]
	for _, value := range args {
		if max < value {
			max = value
		}
	}
	return max
}

func Print_matrix(mtr [][]int) {
	for i := range mtr {
		for j := range mtr[i] {
			fmt.Printf("%3d ", mtr[i][j])
		}
		fmt.Println()
	}
}
