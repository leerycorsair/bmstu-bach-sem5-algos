package ants_test

import (
	"local/ants"
	"testing"
)

func benchmark_brute(cities int, b *testing.B) {
	b.StopTimer()
	citiesMap := ants.GenMap(cities)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ants.BruteForceAlgo(citiesMap, cities)
	}
}

func BenchmarkBrute5(b *testing.B)  { benchmark_brute(2, b) }
func BenchmarkBrute6(b *testing.B)  { benchmark_brute(6, b) }
func BenchmarkBrute7(b *testing.B)  { benchmark_brute(7, b) }
func BenchmarkBrute8(b *testing.B)  { benchmark_brute(8, b) }
func BenchmarkBrute9(b *testing.B)  { benchmark_brute(9, b) }
func BenchmarkBrute10(b *testing.B) { benchmark_brute(10, b) }
func BenchmarkBrute11(b *testing.B) { benchmark_brute(11, b) }
func BenchmarkBrute12(b *testing.B) { benchmark_brute(12, b) }
