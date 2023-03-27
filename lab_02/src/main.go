package main

import (
	"fmt"
	"local/matrix"
)

func menu_print() {
	fmt.Println("\nMenu:")
	fmt.Println("1.Multiply matricies from standart input.")
	fmt.Println("2.Multiply auto-generated matricies.")
	fmt.Println("0.Exit")
}

func menu_option() int {
	var option int
	fmt.Print("Choose an option:")
	fmt.Scanln(&option)
	return option
}

func main() {
	for {
		menu_print()
		option := menu_option()
		switch option {
		case 1:
			fmt.Println("The first matrix:")
			mtr1 := matrix.Read()
			fmt.Println("The second matrix:")
			mtr2 := matrix.Read()
			if matrix.Check_matrix_multi(mtr1, mtr2) {
				fmt.Println("The classic multiplication:")
				res := matrix.Classic_multi(mtr1, mtr2)
				matrix.Print(res)
				fmt.Println("The Winograd multiplication:")
				res = matrix.Winograd_multi(mtr1, mtr2)
				matrix.Print(res)
				fmt.Println("The optimised Winograd multiplication:")
				res = matrix.Winograd_optimised_multi(mtr1, mtr2)
				matrix.Print(res)
			} else {
				fmt.Println("Incorrect dimmensions.")
			}
		case 2:
			fmt.Println("The first matrix:")
			mtr1 := matrix.Gen(0, 0)
			matrix.Print(mtr1)
			fmt.Println("The second matrix:")
			mtr2 := matrix.Gen(0, 0)
			matrix.Print(mtr2)
			if matrix.Check_matrix_multi(mtr1, mtr2) {
				fmt.Println("The classic multiplication:")
				res := matrix.Classic_multi(mtr1, mtr2)
				matrix.Print(res)
				fmt.Println("The Winograd multiplication:")
				res = matrix.Winograd_multi(mtr1, mtr2)
				matrix.Print(res)
				fmt.Println("The optimised Winograd multiplication:")
				res = matrix.Winograd_optimised_multi(mtr1, mtr2)
				matrix.Print(res)
			} else {
				fmt.Println("Incorrect dimmensions.")
			}
		case 0:
			fmt.Println("The program was finished.")
			return
		default:
			fmt.Println("An unknown menu option.")
		}
	}
}
