package levenshtein

import "testing"

func BenchmarkLevRec5(b *testing.B) {
	s1 := "oneor"
	s2 := "minor"

	for i := 0; i < b.N; i++ {
		LevRec(s1, s2)
	}
}

func BenchmarkLevRec8(b *testing.B) {
	s1 := "eightpop"
	s2 := "fruitypo"

	for i := 0; i < b.N; i++ {
		LevRec(s1, s2)
	}
}

func BenchmarkLevRec10(b *testing.B) {
	s1 := "whatswrong"
	s2 := "spitefully"

	for i := 0; i < b.N; i++ {
		LevRec(s1, s2)
	}
}

func BenchmarkLevRec12(b *testing.B) {
	s1 := "aboveeveryth"
	s2 := "whyareyoubad"

	for i := 0; i < b.N; i++ {
		LevRec(s1, s2)
	}
}

// func BenchmarkLevRec15(b *testing.B) {
// 	s1 := "ew6YGS9z40Xhwzp"
// 	s2 := "zDmNyGZ3g4fQgyY"

// 	for i := 0; i < b.N; i++ {
// 		LevRec(s1, s2)
// 	}
// }

// func BenchmarkLevMtr5(b *testing.B) {
// 	s1 := "oneor"
// 	s2 := "minor"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }

// func BenchmarkLevMtr8(b *testing.B) {
// 	s1 := "eightpop"
// 	s2 := "fruitypo"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }

// func BenchmarkLevMtr10(b *testing.B) {
// 	s1 := "whatswrong"
// 	s2 := "spitefully"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }
// func BenchmarkLevMtr12(b *testing.B) {
// 	s1 := "aboveeveryth"
// 	s2 := "whyareyoubad"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }

// func BenchmarkLevMtr20(b *testing.B) {
// 	s1 := "xfvfnhJqzeYEplmC51bP"
// 	s2 := "gO6x1WwiPoQmVEsKZiJb"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }

// func BenchmarkLevMtr40(b *testing.B) {
// 	s1 := "T9VB3welJ7ORQYuES3YKwMoyb5HRM772m6jTsxH8"
// 	s2 := "X9sAXA7n6WCzCAyBCp8wn9PqQJjvW1qZPjSplJdR"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }

// func BenchmarkLevMtr60(b *testing.B) {
// 	s1 := "rpwu7VgA6Ocyc0T4BeUznTyIdwCZS0mHadMPFh0Po4m9Cm17yJ7TVj8qqQ1f"
// 	s2 := "15EFNXiexo0PWXMvcEOVdGbJwDfKCTNmp9yvbYI42kXZSRVrFsuJnt7Dspne"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }

// func BenchmarkLevMtr80(b *testing.B) {
// 	s1 := "1c4ncSlNM6uJh87LpWNZb9gO4lQfitJ6wUnTkxjrcIcBUrt6LCb2uzktBrCowqFf8iWY4RnzQ7Fejbj4"
// 	s2 := "KmJybCrZmNBOWCXupHujRXXyCtk6dPVYzYO3VDiO9jKjZrrJGmmWRFgx5izoOPkWTgQtbOAW1AhQIShU"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }

// func BenchmarkLevMtr100(b *testing.B) {
// 	s1 := "CxGSm4l53FCbxdWKciMijsoDqB1LbkervXkuATWhfq2ok99LUTcQMJTBFf9zqxdORltfWVi5jd1PE4zmdnrI5dMtPpDE88cOpD1n"
// 	s2 := "fsRi5U39GGV0fAsCZLZwIKnwJPLzCH7CpR0AFOSq40RL8siSdBCCyW46aLChc7uTk4epJSjhU5J5uijOQZVKx1nrDwhh8QFU5J80"

// 	for i := 0; i < b.N; i++ {
// 		LevMtr(s1, s2)
// 	}
// }

func BenchmarkLevRecCached5(b *testing.B) {
	s1 := "oneor"
	s2 := "minor"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}

func BenchmarkLevRecCached8(b *testing.B) {
	s1 := "eightpop"
	s2 := "fruitypo"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}

func BenchmarkLevRecCached10(b *testing.B) {
	s1 := "whatswrong"
	s2 := "spitefully"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}
func BenchmarkLevRecCached12(b *testing.B) {
	s1 := "aboveeveryth"
	s2 := "whyareyoubad"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}

func BenchmarkLevRecCached20(b *testing.B) {
	s1 := "xfvfnhJqzeYEplmC51bP"
	s2 := "gO6x1WwiPoQmVEsKZiJb"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}

func BenchmarkLevRecCached40(b *testing.B) {
	s1 := "T9VB3welJ7ORQYuES3YKwMoyb5HRM772m6jTsxH8"
	s2 := "X9sAXA7n6WCzCAyBCp8wn9PqQJjvW1qZPjSplJdR"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}

func BenchmarkLevRecCached60(b *testing.B) {
	s1 := "rpwu7VgA6Ocyc0T4BeUznTyIdwCZS0mHadMPFh0Po4m9Cm17yJ7TVj8qqQ1f"
	s2 := "15EFNXiexo0PWXMvcEOVdGbJwDfKCTNmp9yvbYI42kXZSRVrFsuJnt7Dspne"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}

func BenchmarkLevRecCached80(b *testing.B) {
	s1 := "1c4ncSlNM6uJh87LpWNZb9gO4lQfitJ6wUnTkxjrcIcBUrt6LCb2uzktBrCowqFf8iWY4RnzQ7Fejbj4"
	s2 := "KmJybCrZmNBOWCXupHujRXXyCtk6dPVYzYO3VDiO9jKjZrrJGmmWRFgx5izoOPkWTgQtbOAW1AhQIShU"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}

func BenchmarkLevRecCached100(b *testing.B) {
	s1 := "CxGSm4l53FCbxdWKciMijsoDqB1LbkervXkuATWhfq2ok99LUTcQMJTBFf9zqxdORltfWVi5jd1PE4zmdnrI5dMtPpDE88cOpD1n"
	s2 := "fsRi5U39GGV0fAsCZLZwIKnwJPLzCH7CpR0AFOSq40RL8siSdBCCyW46aLChc7uTk4epJSjhU5J5uijOQZVKx1nrDwhh8QFU5J80"

	for i := 0; i < b.N; i++ {
		LevRecCached(s1, s2)
	}
}

func BenchmarkDamLevMtr5(b *testing.B) {
	s1 := "oneor"
	s2 := "minor"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamLevMtr8(b *testing.B) {
	s1 := "eightpop"
	s2 := "fruitypo"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamLevMtr10(b *testing.B) {
	s1 := "whatswrong"
	s2 := "spitefully"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}
func BenchmarkDamLevMtr12(b *testing.B) {
	s1 := "aboveeveryth"
	s2 := "whyareyoubad"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamLevMtr20(b *testing.B) {
	s1 := "xfvfnhJqzeYEplmC51bP"
	s2 := "gO6x1WwiPoQmVEsKZiJb"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamLevMtr40(b *testing.B) {
	s1 := "T9VB3welJ7ORQYuES3YKwMoyb5HRM772m6jTsxH8"
	s2 := "X9sAXA7n6WCzCAyBCp8wn9PqQJjvW1qZPjSplJdR"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamLevMtr60(b *testing.B) {
	s1 := "rpwu7VgA6Ocyc0T4BeUznTyIdwCZS0mHadMPFh0Po4m9Cm17yJ7TVj8qqQ1f"
	s2 := "15EFNXiexo0PWXMvcEOVdGbJwDfKCTNmp9yvbYI42kXZSRVrFsuJnt7Dspne"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamLevMtr80(b *testing.B) {
	s1 := "1c4ncSlNM6uJh87LpWNZb9gO4lQfitJ6wUnTkxjrcIcBUrt6LCb2uzktBrCowqFf8iWY4RnzQ7Fejbj4"
	s2 := "KmJybCrZmNBOWCXupHujRXXyCtk6dPVYzYO3VDiO9jKjZrrJGmmWRFgx5izoOPkWTgQtbOAW1AhQIShU"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamLevMtr100(b *testing.B) {
	s1 := "CxGSm4l53FCbxdWKciMijsoDqB1LbkervXkuATWhfq2ok99LUTcQMJTBFf9zqxdORltfWVi5jd1PE4zmdnrI5dMtPpDE88cOpD1n"
	s2 := "fsRi5U39GGV0fAsCZLZwIKnwJPLzCH7CpR0AFOSq40RL8siSdBCCyW46aLChc7uTk4epJSjhU5J5uijOQZVKx1nrDwhh8QFU5J80"

	for i := 0; i < b.N; i++ {
		DamLevMtr(s1, s2)
	}
}

func BenchmarkDamLevRec5(b *testing.B) {
	s1 := "oneor"
	s2 := "minor"

	for i := 0; i < b.N; i++ {
		DamLevRec(s1, s2)
	}
}

func BenchmarkDamLevRec8(b *testing.B) {
	s1 := "eightpop"
	s2 := "fruitypo"

	for i := 0; i < b.N; i++ {
		DamLevRec(s1, s2)
	}
}

func BenchmarkDamLevRec10(b *testing.B) {
	s1 := "whatswrong"
	s2 := "spitefully"

	for i := 0; i < b.N; i++ {
		DamLevRec(s1, s2)
	}
}
func BenchmarkDamLevRec12(b *testing.B) {
	s1 := "aboveeveryth"
	s2 := "whyareyoubad"

	for i := 0; i < b.N; i++ {
		DamLevRec(s1, s2)
	}
}

// func BenchmarkDamLevRec15(b *testing.B) {
// 	s1 := "ew6YGS9z40Xhwzp"
// 	s2 := "zDmNyGZ3g4fQgyY"

// 	for i := 0; i < b.N; i++ {
// 		DamLevRec(s1, s2)
// 	}
// }
