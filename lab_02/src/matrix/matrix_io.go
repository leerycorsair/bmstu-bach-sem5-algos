package matrix

import "fmt"

func Print(mtr MatrixInt) {
	for i := 0; i < mtr.rows; i++ {
		for j := 0; j < mtr.columns; j++ {
			fmt.Printf("%4d ", mtr.values[i][j])
		}
		fmt.Println()
	}
}

func Read() MatrixInt {
	var rows int
	fmt.Print("Rows:")
	fmt.Scanln(&rows)
	var columns int
	fmt.Print("Columns:")
	fmt.Scanln(&columns)
	mtr := matrix_init(rows, columns)

	for i := 0; i < mtr.rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Scanf("%d", &mtr.values[i][j])
		}
	}
	return mtr
}

func matrix_init(rows int, colunms int) MatrixInt {
	var mtr MatrixInt
	mtr.rows = rows
	mtr.columns = colunms
	mtr.values = make([][]int, mtr.rows)
	for i := range mtr.values {
		mtr.values[i] = make([]int, mtr.columns)
	}
	return mtr
}
