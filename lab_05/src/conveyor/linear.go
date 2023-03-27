package conveyor

import "time"

func Linear(orders int) []*Ration_t {
	var secondLineQ, thirdLineQ Queue
	var result []*Ration_t

	for {
		if len(result) != orders {
			ration := new(Ration_t)

			ration.startBreakfast = time.Now()
			ration.breakfast = genBreakfast()
			ration.finishBreakfast = time.Now()

			secondLineQ.push(ration)

			if !secondLineQ.isEmpty() {
				ration = secondLineQ.pop()

				ration.startLunch = time.Now()
				ration.lunch = genLunch()
				ration.finishLunch = time.Now()

				thirdLineQ.push(ration)
			}

			if !thirdLineQ.isEmpty() {
				ration = thirdLineQ.pop()

				ration.startDinner = time.Now()
				ration.dinner = genDinner()
				ration.finishDinner = time.Now()

				result = append(result, ration)
			}
		} else {
			return result
		}
	}
}
