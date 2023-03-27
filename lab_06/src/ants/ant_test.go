package ants_test

import (
	"local/ants"
	"testing"
)

func benchmark_ants(cities int, b *testing.B) {
	b.StopTimer()
	citiesMap := ants.GenMap(cities)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		ants.AntAlgo(citiesMap, ants.DEFAULT_ALPHA, ants.DEFAULT_BETA, ants.DEFAULT_P, ants.DEFAULT_PHEROMONCOEF, ants.DEFAULT_DAYS)
	}
}

func BenchmarkAnts5(b *testing.B)  { benchmark_ants(2, b) }
func BenchmarkAnts6(b *testing.B)  { benchmark_ants(6, b) }
func BenchmarkAnts7(b *testing.B)  { benchmark_ants(7, b) }
func BenchmarkAnts8(b *testing.B)  { benchmark_ants(8, b) }
func BenchmarkAnts9(b *testing.B)  { benchmark_ants(9, b) }
func BenchmarkAnts10(b *testing.B) { benchmark_ants(10, b) }
func BenchmarkAnts11(b *testing.B) { benchmark_ants(11, b) }
func BenchmarkAnts12(b *testing.B) { benchmark_ants(12, b) }
