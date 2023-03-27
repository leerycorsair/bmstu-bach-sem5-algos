package matrix

import "testing"

func benchmark_multi(size int, b *testing.B, f func(MatrixInt, MatrixInt) MatrixInt) {
	b.StopTimer()
	mtr1 := Gen(size, size)
	mtr2 := Gen(size, size)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		f(mtr1, mtr2)
	}
}

func BenchmarkCM100(b *testing.B) { benchmark_multi(100, b, Classic_multi) }
func BenchmarkCM200(b *testing.B) { benchmark_multi(200, b, Classic_multi) }
func BenchmarkCM300(b *testing.B) { benchmark_multi(300, b, Classic_multi) }
func BenchmarkCM400(b *testing.B) { benchmark_multi(400, b, Classic_multi) }
func BenchmarkCM500(b *testing.B) { benchmark_multi(500, b, Classic_multi) }
func BenchmarkCM600(b *testing.B) { benchmark_multi(600, b, Classic_multi) }
func BenchmarkCM700(b *testing.B) { benchmark_multi(700, b, Classic_multi) }

// func BenchmarkCM800(b *testing.B)  { benchmark_multi(800, b, Classic_multi) }
// func BenchmarkCM900(b *testing.B)  { benchmark_multi(900, b, Classic_multi) }
// func BenchmarkCM1000(b *testing.B) { benchmark_multi(1000, b, Classic_multi) }

func BenchmarkCM101(b *testing.B) { benchmark_multi(101, b, Classic_multi) }
func BenchmarkCM201(b *testing.B) { benchmark_multi(201, b, Classic_multi) }
func BenchmarkCM301(b *testing.B) { benchmark_multi(301, b, Classic_multi) }
func BenchmarkCM401(b *testing.B) { benchmark_multi(401, b, Classic_multi) }
func BenchmarkCM501(b *testing.B) { benchmark_multi(501, b, Classic_multi) }
func BenchmarkCM601(b *testing.B) { benchmark_multi(601, b, Classic_multi) }
func BenchmarkCM701(b *testing.B) { benchmark_multi(701, b, Classic_multi) }

// func BenchmarkCM801(b *testing.B)  { benchmark_multi(801, b, Classic_multi) }
// func BenchmarkCM901(b *testing.B)  { benchmark_multi(901, b, Classic_multi) }
// func BenchmarkCM1001(b *testing.B) { benchmark_multi(1001, b, Classic_multi) }

func BenchmarkWM100(b *testing.B) { benchmark_multi(100, b, Winograd_multi) }
func BenchmarkWM200(b *testing.B) { benchmark_multi(200, b, Winograd_multi) }
func BenchmarkWM300(b *testing.B) { benchmark_multi(300, b, Winograd_multi) }
func BenchmarkWM400(b *testing.B) { benchmark_multi(400, b, Winograd_multi) }
func BenchmarkWM500(b *testing.B) { benchmark_multi(500, b, Winograd_multi) }
func BenchmarkWM600(b *testing.B) { benchmark_multi(600, b, Winograd_multi) }
func BenchmarkWM700(b *testing.B) { benchmark_multi(700, b, Winograd_multi) }

// func BenchmarkWM800(b *testing.B)  { benchmark_multi(800, b, Winograd_multi) }
// func BenchmarkWM900(b *testing.B)  { benchmark_multi(900, b, Winograd_multi) }
// func BenchmarkWM1000(b *testing.B) { benchmark_multi(1000, b, Winograd_multi) }

func BenchmarkWM101(b *testing.B) { benchmark_multi(101, b, Winograd_multi) }
func BenchmarkWM201(b *testing.B) { benchmark_multi(201, b, Winograd_multi) }
func BenchmarkWM301(b *testing.B) { benchmark_multi(301, b, Winograd_multi) }
func BenchmarkWM401(b *testing.B) { benchmark_multi(401, b, Winograd_multi) }
func BenchmarkWM501(b *testing.B) { benchmark_multi(501, b, Winograd_multi) }
func BenchmarkWM601(b *testing.B) { benchmark_multi(601, b, Winograd_multi) }
func BenchmarkWM701(b *testing.B) { benchmark_multi(701, b, Winograd_multi) }

// func BenchmarkWM801(b *testing.B)  { benchmark_multi(801, b, Winograd_multi) }
// func BenchmarkWM901(b *testing.B)  { benchmark_multi(901, b, Winograd_multi) }
// func BenchmarkWM1001(b *testing.B) { benchmark_multi(1001, b, Winograd_multi) }

func BenchmarkWMO100(b *testing.B) { benchmark_multi(100, b, Winograd_optimised_multi) }
func BenchmarkWMO200(b *testing.B) { benchmark_multi(200, b, Winograd_optimised_multi) }
func BenchmarkWMO300(b *testing.B) { benchmark_multi(300, b, Winograd_optimised_multi) }
func BenchmarkWMO400(b *testing.B) { benchmark_multi(400, b, Winograd_optimised_multi) }
func BenchmarkWMO500(b *testing.B) { benchmark_multi(500, b, Winograd_optimised_multi) }
func BenchmarkWMO600(b *testing.B) { benchmark_multi(600, b, Winograd_optimised_multi) }
func BenchmarkWMO700(b *testing.B) { benchmark_multi(700, b, Winograd_optimised_multi) }

// func BenchmarkWMO800(b *testing.B)  { benchmark_multi(800, b, Winograd_optimised_multi) }
// func BenchmarkWMO900(b *testing.B)  { benchmark_multi(900, b, Winograd_optimised_multi) }
// func BenchmarkWMO1000(b *testing.B) { benchmark_multi(1000, b, Winograd_optimised_multi) }

func BenchmarkWMO101(b *testing.B) { benchmark_multi(101, b, Winograd_optimised_multi) }
func BenchmarkWMO201(b *testing.B) { benchmark_multi(201, b, Winograd_optimised_multi) }
func BenchmarkWMO301(b *testing.B) { benchmark_multi(301, b, Winograd_optimised_multi) }
func BenchmarkWMO401(b *testing.B) { benchmark_multi(401, b, Winograd_optimised_multi) }
func BenchmarkWMO501(b *testing.B) { benchmark_multi(501, b, Winograd_optimised_multi) }
func BenchmarkWMO601(b *testing.B) { benchmark_multi(601, b, Winograd_optimised_multi) }
func BenchmarkWMO701(b *testing.B) { benchmark_multi(701, b, Winograd_optimised_multi) }

// func BenchmarkWMO801(b *testing.B)  { benchmark_multi(801, b, Winograd_optimised_multi) }
// func BenchmarkWMO901(b *testing.B)  { benchmark_multi(901, b, Winograd_optimised_multi) }
// func BenchmarkWMO1001(b *testing.B) { benchmark_multi(1001, b, Winograd_optimised_multi) }
