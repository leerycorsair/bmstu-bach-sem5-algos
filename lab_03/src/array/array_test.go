package array_test

import (
	"local/array"
	"testing"
)

// func BenchmarkBS100rand(b *testing.B) {
// 	arr := array.Arr_generate(100, "rand")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_bubble_sort(arrCopy)
// 	}
// }

func BenchmarkBS500rand(b *testing.B) {
	arr := array.Arr_generate(500, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS1000rand(b *testing.B) {
	arr := array.Arr_generate(1000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS2000rand(b *testing.B) {
	arr := array.Arr_generate(2000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS3000rand(b *testing.B) {
	arr := array.Arr_generate(3000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS5000rand(b *testing.B) {
	arr := array.Arr_generate(5000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

// func BenchmarkBS100asc(b *testing.B) {
// 	arr := array.Arr_generate(100, "asc")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_bubble_sort(arrCopy)
// 	}
// }

func BenchmarkBS500asc(b *testing.B) {
	arr := array.Arr_generate(500, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS1000asc(b *testing.B) {
	arr := array.Arr_generate(1000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS2000asc(b *testing.B) {
	arr := array.Arr_generate(2000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS3000asc(b *testing.B) {
	arr := array.Arr_generate(3000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS5000asc(b *testing.B) {
	arr := array.Arr_generate(5000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

// func BenchmarkBS100desc(b *testing.B) {
// 	arr := array.Arr_generate(100, "desc")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_bubble_sort(arrCopy)
// 	}
// }

func BenchmarkBS500desc(b *testing.B) {
	arr := array.Arr_generate(500, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS1000desc(b *testing.B) {
	arr := array.Arr_generate(1000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS2000desc(b *testing.B) {
	arr := array.Arr_generate(2000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS3000desc(b *testing.B) {
	arr := array.Arr_generate(3000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

func BenchmarkBS5000desc(b *testing.B) {
	arr := array.Arr_generate(5000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_bubble_sort(arrCopy)
	}
}

// func BenchmarkIS100rand(b *testing.B) {
// 	arr := array.Arr_generate(100, "rand")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_insertion_sort(arrCopy)
// 	}
// }

func BenchmarkIS500rand(b *testing.B) {
	arr := array.Arr_generate(500, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS1000rand(b *testing.B) {
	arr := array.Arr_generate(1000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS2000rand(b *testing.B) {
	arr := array.Arr_generate(2000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS3000rand(b *testing.B) {
	arr := array.Arr_generate(3000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS5000rand(b *testing.B) {
	arr := array.Arr_generate(5000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

// func BenchmarkIS100asc(b *testing.B) {
// 	arr := array.Arr_generate(100, "asc")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_insertion_sort(arrCopy)
// 	}
// }

func BenchmarkIS500asc(b *testing.B) {
	arr := array.Arr_generate(500, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS1000asc(b *testing.B) {
	arr := array.Arr_generate(1000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS2000asc(b *testing.B) {
	arr := array.Arr_generate(2000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS3000asc(b *testing.B) {
	arr := array.Arr_generate(3000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS5000asc(b *testing.B) {
	arr := array.Arr_generate(5000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

// func BenchmarkIS100desc(b *testing.B) {
// 	arr := array.Arr_generate(100, "desc")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_insertion_sort(arrCopy)
// 	}
// }

func BenchmarkIS500desc(b *testing.B) {
	arr := array.Arr_generate(500, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS1000desc(b *testing.B) {
	arr := array.Arr_generate(1000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS2000desc(b *testing.B) {
	arr := array.Arr_generate(2000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS3000desc(b *testing.B) {
	arr := array.Arr_generate(3000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

func BenchmarkIS5000desc(b *testing.B) {
	arr := array.Arr_generate(5000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_insertion_sort(arrCopy)
	}
}

// func BenchmarkSS100rand(b *testing.B) {
// 	arr := array.Arr_generate(100, "rand")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_selection_sort(arrCopy)
// 	}
// }

func BenchmarkSS500rand(b *testing.B) {
	arr := array.Arr_generate(500, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS1000rand(b *testing.B) {
	arr := array.Arr_generate(1000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS2000rand(b *testing.B) {
	arr := array.Arr_generate(2000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS3000rand(b *testing.B) {
	arr := array.Arr_generate(3000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS5000rand(b *testing.B) {
	arr := array.Arr_generate(5000, "rand")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

// func BenchmarkSS100asc(b *testing.B) {
// 	arr := array.Arr_generate(100, "asc")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_selection_sort(arrCopy)
// 	}
// }

func BenchmarkSS500asc(b *testing.B) {
	arr := array.Arr_generate(500, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS1000asc(b *testing.B) {
	arr := array.Arr_generate(1000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS2000asc(b *testing.B) {
	arr := array.Arr_generate(2000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS3000asc(b *testing.B) {
	arr := array.Arr_generate(3000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS5000asc(b *testing.B) {
	arr := array.Arr_generate(5000, "asc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

// func BenchmarkSS100desc(b *testing.B) {
// 	arr := array.Arr_generate(100, "desc")

// 	for i := 0; i < b.N; i++ {
// 		b.StopTimer()
// 		arrCopy := make([]int, len(arr))
// 		copy(arrCopy, arr)
// 		b.StartTimer()
// 		array.Arr_selection_sort(arrCopy)
// 	}
// }

func BenchmarkSS500desc(b *testing.B) {
	arr := array.Arr_generate(500, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS1000desc(b *testing.B) {
	arr := array.Arr_generate(1000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS2000desc(b *testing.B) {
	arr := array.Arr_generate(2000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS3000desc(b *testing.B) {
	arr := array.Arr_generate(3000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}

func BenchmarkSS5000desc(b *testing.B) {
	arr := array.Arr_generate(5000, "desc")

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrCopy := make([]int, len(arr))
		copy(arrCopy, arr)
		b.StartTimer()
		array.Arr_selection_sort(arrCopy)
	}
}
