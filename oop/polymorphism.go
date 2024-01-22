// polymorphism is about defining objects that change behavior based on context
//go acheives polymorphism through interfaces
// a type implements an interface by providing funciton definitions for each method declared within the interface
// different types can implement the interface in different ways

package main

import (
	"fmt"
)

type Menu interface {
	MakeBreakfast() string
	MakeLunch() string
	MakeDinner() string
}

type Restaurant struct {
	BestBreakfast string
	BestLunch     string
	BestDinner    string
}

// implement Menu interface for type Restaurant
func (r Restaurant) MakeBreakfast() string {
	return r.BestBreakfast
}

func (r Restaurant) MakeLunch() string {
	return r.BestLunch
}

func (r Restaurant) MakeDinner() string {
	return r.BestDinner
}

func cook_all(m Menu) {
	fmt.Println("Breakfast:", m.MakeBreakfast())
	fmt.Println("Lunch:", m.MakeLunch())
	fmt.Println("Dinner:", m.MakeDinner())
}

func display_food() {
	gordon_ramsay_hells_kitchen := Restaurant{
		BestBreakfast: "steak&eggs",
		BestLunch:     "oysters",
		BestDinner:    "king crab",
	}
	cook_all(gordon_ramsay_hells_kitchen)
}
