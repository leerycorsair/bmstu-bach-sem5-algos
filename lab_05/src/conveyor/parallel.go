package conveyor

import (
	"sync"
	"time"
)

func Parallel(orders int) []*Ration_t {
	firstLine := make(chan *Ration_t, orders)
	secondLine := make(chan *Ration_t, orders)
	thirdLine := make(chan *Ration_t, orders)
	result := make([]*Ration_t, 0, orders)
	var wg sync.WaitGroup
	wg.Add(3)
	breakfast := func() {
		for ration := range firstLine {
			ration.startBreakfast = time.Now()
			ration.breakfast = genBreakfast()
			ration.finishBreakfast = time.Now()
			secondLine <- ration
		}
	}
	lunch := func() {
		for ration := range secondLine {
			ration.startLunch = time.Now()
			ration.lunch = genLunch()
			ration.finishLunch = time.Now()
			thirdLine <- ration
		}
	}
	dinner := func() {
		for ration := range thirdLine {
			ration.startDinner = time.Now()
			ration.dinner = genDinner()
			ration.finishDinner = time.Now()
			result = append(result, ration)

			if len(result) == cap(result) {
				wg.Done()
				wg.Done()
				wg.Done()
			}
		}
	}
	go breakfast()
	go lunch()
	go dinner()
	for i := 0; i < orders; i++ {
		ration := new(Ration_t)
		firstLine <- ration
	}
	wg.Wait()
	return result
}
