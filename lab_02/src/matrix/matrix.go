package matrix

import (
	"fmt"
	"math/rand"
)

func Gen(rows int, columns int) MatrixInt {
	if rows == 0 && columns == 0 {
		fmt.Print("Rows:")
		fmt.Scanln(&rows)
		fmt.Print("Columns:")
		fmt.Scanln(&columns)
	}
	mtr := matrix_init(rows, columns)

	for i := 0; i < mtr.rows; i++ {
		for j := 0; j < columns; j++ {
			mtr.values[i][j] = rand.Intn(100)
		}
	}
	return mtr
}

func Check_matrix_multi(mtr1 MatrixInt, mtr2 MatrixInt) bool {
	return !(mtr1.columns != mtr2.rows)
}

func Classic_multi(mtr1 MatrixInt, mtr2 MatrixInt) MatrixInt {
	res := matrix_init(mtr1.rows, mtr2.columns)
	for i := 0; i < mtr1.rows; i++ {
		for j := 0; j < mtr2.columns; j++ {
			for k := 0; k < mtr1.columns; k++ {
				res.values[i][j] += mtr1.values[i][k] * mtr2.values[k][j]
			}
		}
	}
	return res
}

func Winograd_multi(mtr1 MatrixInt, mtr2 MatrixInt) MatrixInt {
	res := matrix_init(mtr1.rows, mtr2.columns)
	wrowcf := calc_wrows(mtr1)
	wcolumncf := calc_wcolumns(mtr2)
	for i := 0; i < res.rows; i++ {
		for j := 0; j < res.columns; j++ {
			res.values[i][j] = -wrowcf[i] - wcolumncf[j]
			for k := 0; k < mtr2.rows/2; k++ {
				res.values[i][j] = res.values[i][j] + (mtr1.values[i][k*2]+mtr2.values[k*2+1][j])*(mtr1.values[i][k*2+1]+
					mtr2.values[k*2][j])
			}
		}
	}
	if mtr2.rows%2 != 0 {
		for i := 0; i < res.rows; i++ {
			for j := 0; j < res.columns; j++ {
				res.values[i][j] += mtr1.values[i][mtr1.columns-1] * mtr2.values[mtr2.rows-1][j]
			}
		}
	}
	return res
}

func calc_wrows(mtr MatrixInt) []int {
	wrowcf := make([]int, mtr.rows)
	for i := 0; i < mtr.rows; i++ {
		for j := 0; j < mtr.columns/2; j++ {
			wrowcf[i] += mtr.values[i][j*2] * mtr.values[i][j*2+1]
		}
	}
	return wrowcf
}

func calc_wcolumns(mtr MatrixInt) []int {
	wcolumncf := make([]int, mtr.columns)
	for i := 0; i < mtr.columns; i++ {
		for j := 0; j < mtr.rows/2; j++ {
			wcolumncf[i] += mtr.values[j*2][i] * mtr.values[j*2+1][i]
		}
	}
	return wcolumncf
}

func Winograd_optimised_multi(mtr1 MatrixInt, mtr2 MatrixInt) MatrixInt {
	res := matrix_init(mtr1.rows, mtr2.columns)
	wrowcf := calc_wrows_optimised(mtr1)
	wcolumncf := calc_wcolumns_optimised(mtr2)
	for i := 0; i < res.rows; i++ {
		for j := 0; j < res.columns; j++ {
			buf := 0
			for k := 0; k < mtr2.rows/2; k++ {
				buf += (mtr1.values[i][k*2] + mtr2.values[k*2+1][j]) * (mtr1.values[i][k*2+1] +
					mtr2.values[k*2][j])
			}
			res.values[i][j] = buf - wrowcf[i] - wcolumncf[j]
		}
	}
	if mtr2.rows%2 != 0 {
		k := mtr1.columns - 1
		for i := 0; i < res.rows; i++ {
			for j := 0; j < res.columns; j++ {
				res.values[i][j] += mtr1.values[i][k] * mtr2.values[k][j]
			}
		}
	}
	return res
}
func calc_wrows_optimised(mtr MatrixInt) []int {
	wrowcf := make([]int, mtr.rows)
	for i := 0; i < mtr.rows; i++ {
		buf := 0
		for j := 0; j < mtr.columns/2; j++ {
			buf += mtr.values[i][j*2] * mtr.values[i][j*2+1]
		}
		wrowcf[i] = buf
	}
	return wrowcf
}
func calc_wcolumns_optimised(mtr MatrixInt) []int {
	wcolumncf := make([]int, mtr.columns)
	for i := 0; i < mtr.columns; i++ {
		buf := 0
		for j := 0; j < mtr.rows/2; j++ {
			buf += mtr.values[j*2][i] * mtr.values[j*2+1][i]
		}
		wcolumncf[i] = buf
	}
	return wcolumncf
}
