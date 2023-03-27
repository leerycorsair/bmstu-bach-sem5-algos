package conveyor

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogTime(orders []*Ration_t) {
	timestamp := orders[0].startBreakfast

	fmt.Printf("Stage Starting Time\n")
	for id, order := range orders {
		fmt.Printf("Order ID = %3v, 1 Stage = %v, 2 Stage = %v, 3 Stage = %v\n",
			id,
			order.startBreakfast.Sub(timestamp),
			order.startLunch.Sub(timestamp),
			order.startDinner.Sub(timestamp))
	}

	fmt.Printf("Stage Finishing Time\n")
	for id, order := range orders {
		fmt.Printf("Order ID = %3v, 1 Stage = %v, 2 Stage = %v, 3 Stage = %v\n",
			id,
			order.finishBreakfast.Sub(timestamp),
			order.finishLunch.Sub(timestamp),
			order.finishDinner.Sub(timestamp))
	}

	var fDowntime, sDowntime, tDowntime time.Duration
	for i := 1; i < len(orders); i++ {
		fDowntime += orders[i].startBreakfast.Sub(timestamp) - orders[i-1].finishBreakfast.Sub(timestamp)
		sDowntime += orders[i].startLunch.Sub(timestamp) - orders[i-1].finishLunch.Sub(timestamp)
		tDowntime += orders[i].startDinner.Sub(timestamp) - orders[i-1].finishDinner.Sub(timestamp)
	}
	fmt.Printf("Downtimes\n")
	fmt.Printf("    1 Stage = %v, 2 Stage = %v, 3 Stage = %v\n",
		fDowntime, sDowntime, tDowntime)
}

func LogTimeLatex(orders []*Ration_t, filename string) {
	timestamp := orders[0].startBreakfast

	file, err := os.Create("./" + filename + ".csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "order_id,stage_id,time_start,time_end\n")
	for id, order := range orders {
		fmt.Fprintf(file, " %v&%v&%v&%v\\\\ ", id+1, 1, order.startBreakfast.Sub(timestamp), order.finishBreakfast.Sub(timestamp))
		fmt.Fprintf(file, " %v&%v&%v&%v\\\\ ", id+1, 2, order.startLunch.Sub(timestamp), order.finishLunch.Sub(timestamp))
		fmt.Fprintf(file, " %v&%v&%v&%v\\\\ \\hline", id+1, 3, order.startDinner.Sub(timestamp), order.finishDinner.Sub(timestamp))
	}
}

func LogData(orders []*Ration_t) {
	for _, order := range orders {
		fmt.Printf("Breakfast: %v and %v\n", order.breakfast.maincourse, order.breakfast.fruit)
		fmt.Printf("Lunch: %v and %v\n", order.lunch.maincourse, order.lunch.snack)
		fmt.Printf("Dinner: %v and %v\n\n\n", order.dinner.maincourse, order.dinner.dessert)
	}
}
