package main

import (
	"fmt"
	"local/array"
)

func menu_print() {
	fmt.Println("\nMenu:")
	fmt.Println("1.Sort an array from standart input.")
	fmt.Println("2.Sort auto-generated arrays.")
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
			arr := array.Arr_read()

			fmt.Println("\nThe original array:")
			array.Arr_print(arr)
			arrCopy := make([]int, len(arr))

			fmt.Println("Bubble sort:")
			copy(arrCopy, arr)
			array.Arr_bubble_sort(arrCopy)
			array.Arr_print(arrCopy)

			fmt.Println("Insertion sort:")
			copy(arrCopy, arr)
			array.Arr_insertion_sort(arrCopy)
			array.Arr_print(arrCopy)

			fmt.Println("Selection sort:")
			copy(arrCopy, arr)
			array.Arr_selection_sort(arrCopy)
			array.Arr_print(arrCopy)
		case 2:
			var size int
			fmt.Println("Enter an array size:")
			fmt.Scanln(&size)

			arr := array.Arr_generate(size, "rand")

			fmt.Println("\nThe original array:")
			array.Arr_print(arr)
			arrCopy := make([]int, len(arr))

			fmt.Println("Bubble sort:")
			copy(arrCopy, arr)
			array.Arr_bubble_sort(arrCopy)
			array.Arr_print(arrCopy)

			fmt.Println("Insertion sort:")
			copy(arrCopy, arr)
			array.Arr_insertion_sort(arrCopy)
			array.Arr_print(arrCopy)

			fmt.Println("Selection sort:")
			copy(arrCopy, arr)
			array.Arr_selection_sort(arrCopy)
			array.Arr_print(arrCopy)
		case 0:
			fmt.Println("The program was finished.")
			return
		default:
			fmt.Println("An unknown menu option.")
		}
	}
}
