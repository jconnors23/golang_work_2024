// go acheives polymorphism through interfaces
// a type implements an interface by providing funciton definitions for each method declared within the interface
// different types can implement the interface in different ways

package main

import (
	"fmt"
)

type Menu interface {
	make_breakfast() string
	make_lunch() string
	make_dinner() string
}

type Restuarant struct {
	best_breakfast string
	best_lunch     string
	best_dinner    string
}

func (r Restuarant) make_breakfast() string {
	return r.best_breakfast
}

func (r Restuarant) make_lunch() string {
	return r.best_lunch
}

func (r Restuarant) make_dinner() string {
	return r.best_dinner
}

func cook_all(r Restuarant) {
	fmt.Println("Breakfast:", r.make_breakfast())
	fmt.Println("Lunch:", r.make_lunch())
	fmt.Println("Dinner:", r.make_dinner())
}

func display_food() {
	gordon_ramsay_hells_kitchen := Restuarant{
		best_breakfast: "steak&eggs",
		best_lunch:     "oysters",
		best_dinner:    "king crab",
	}
	fmt.Println(cook_all(gordon_ramsay_hells_kitchen))
}
