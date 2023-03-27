package conveyor

import "time"

type Ration_t struct {
	startBreakfast  time.Time
	finishBreakfast time.Time
	startLunch      time.Time
	finishLunch     time.Time
	startDinner     time.Time
	finishDinner    time.Time

	breakfast Breakfast_t
	lunch     Lunch_t
	dinner    Dinner_t
}

type Breakfast_t struct {
	maincourse string
	fruit      string
}

type Lunch_t struct {
	maincourse string
	snack      string
}

type Dinner_t struct {
	maincourse string
	dessert    string
}
