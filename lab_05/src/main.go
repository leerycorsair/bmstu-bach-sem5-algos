package main

import (
	"fmt"
	"local/conveyor"
	"log"
	"math/rand"
	"os"
	"runtime"
	"time"
)

func checkFuncTime(f func(int) []*conveyor.Ration_t, arg int) time.Duration {
	start := time.Now()
	f(arg)
	end := time.Now()
	return end.Sub(start)
}

func analysis(min, max, step int) {
	fileParallel, err := os.Create("./parallel.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileParallel.Close()

	fileLinear, err := os.Create("./linear.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileLinear.Close()

	fmt.Fprintf(fileParallel, "recs,time\n")
	fmt.Fprintf(fileLinear, "recs,time\n")

	for i := min; i <= max; i += step {
		PTime := checkFuncTime(conveyor.Parallel, i).Milliseconds()
		LTime := checkFuncTime(conveyor.Linear, i).Milliseconds()

		fmt.Printf("Orders = %v, PTime = %v, LTime = %v\n", i, PTime, LTime)
		fmt.Fprintf(fileParallel, "%v,%v\n", i, PTime)
		fmt.Fprintf(fileLinear, "%v,%v\n", i, LTime)
	}
}

func main() {
	runtime.GOMAXPROCS(12)
	rand.Seed(time.Now().UTC().UnixNano())

	var orders int
	fmt.Print("Enter the amount of orders:")
	fmt.Scanln(&orders)

	fmt.Println("Parallel algorithm:")
	result := conveyor.Parallel(orders)
	conveyor.LogTime(result)
	conveyor.LogTimeLatex(result, "p_time")

	fmt.Println("Linear algorithm:")
	result = conveyor.Linear(orders)
	conveyor.LogTime(result)
	conveyor.LogTimeLatex(result, "l_time")

	// conveyor.LogData(result)

	analysis(40, 400, 40)
}
