package array

import (
	"fmt"
	"math/rand"
)

func Arr_read() []int {
	var size int
	fmt.Println("Enter an array size:")
	fmt.Scanln(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	return arr
}

func Arr_print(arr []int) {
	fmt.Printf("%v\n", arr)
}

func Arr_bubble_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		swap := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}

func Arr_insertion_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		elem := arr[i]
		j := i - 1
		for ; j >= 0 && elem < arr[j]; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = elem
	}
}

func Arr_selection_sort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min_elem := arr[i]
		min_index := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min_elem {
				min_elem = arr[j]
				min_index = j
			}
		}
		arr[min_index], arr[i] = arr[i], arr[min_index]
	}
}

func Arr_generate(size int, gen_type string) []int {
	arr := make([]int, size)
	switch gen_type {
	case "asc":
		arr[0] = rand.Intn(50)
		for i := 1; i < size; i++ {
			arr[i] = arr[i-1] + rand.Intn(50)
		}
	case "desc":
		arr[0] = rand.Intn(50)
		for i := 1; i < size; i++ {
			arr[i] = arr[i-1] - rand.Intn(50)
		}
	case "rand":
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(50)
		}
	}
	return arr
}
