package main

import (
	"fmt"
	"math/rand"
	"time"
)

func matrix_generate(size int) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]float64, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = float64(rand.Intn(20)) + 1
		}
	}
	return matrix
}

func matrix_deepcopy(matrix [][]float64) [][]float64 {
	matrix_copy := make([][]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		matrix_copy[i] = make([]float64, len(matrix))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix_copy[i][j] = matrix[i][j]
		}
	}
	return matrix_copy
}

func matrix_print(matrix [][]float64) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%5.3v ", matrix[i][j])
		}
		fmt.Println()
	}
}

func main() {
	size := 3                             //1
	Matrix := matrix_generate(size)       // 2
	Buf_Matrix := matrix_deepcopy(Matrix) // 3

	for step := 0; step < size-1; step++ { // 4
		for row := step + 1; row < size; row++ { // 5
			coeff := -Buf_Matrix[row][step] / Buf_Matrix[step][step] // 6
			for col := step; col < size; col++ {                     // 7
				Buf_Matrix[row][col] += Buf_Matrix[step][col] * coeff // 8
			}
		}
	}

	det := 1.0                  // 9
	for i := 0; i < size; i++ { // 10
		det *= Buf_Matrix[i][i] // 11
	}

	fmt.Println("The original matrix")         // 12
	matrix_print(Matrix)                       // 13
	fmt.Println("The upper triangular matrix") // 14
	matrix_print(Buf_Matrix)                   // 15
	fmt.Printf("Determinant =  %v\n", det)     // 16
}
