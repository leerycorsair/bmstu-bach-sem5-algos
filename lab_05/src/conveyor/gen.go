package conveyor

import (
	"math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func genBreakfast() Breakfast_t {
	breakfast := new(Breakfast_t)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
	breakfast.maincourse = gofakeit.Breakfast()
	breakfast.fruit = gofakeit.Fruit()
	return *breakfast
}

func genLunch() Lunch_t {
	lunch := new(Lunch_t)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
	lunch.maincourse = gofakeit.Lunch()
	lunch.snack = gofakeit.Snack()
	return *lunch
}

func genDinner() Dinner_t {
	dinner := new(Dinner_t)
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(20)))
	dinner.maincourse = gofakeit.Dinner()
	dinner.dessert = gofakeit.Dessert()
	return *dinner
}
