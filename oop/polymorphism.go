// polymorphism is about defining objects that change behavior based on context
//go acheives polymorphism through interfaces
// a type implements an interface by providing funciton definitions for each method declared within the interface
// different types can implement the interface in different ways

package main

import (
	"fmt"
)

type Menu interface {
	makeBreakfast() string
	makeLunch() string
	makeDinner() string
}

type Restaurant struct {
	BestBreakfast string
	BestLunch     string
	BestDinner    string
}

// implement Menu interface for type Restaurant
func (r Restaurant) makeBreakfast() string {
	return r.BestBreakfast
}

func (r Restaurant) makeLunch() string {
	return r.BestLunch
}

func (r Restaurant) makeDinner() string {
	return r.BestDinner
}

func cookAll(m Menu) {
	fmt.Println("Breakfast:", m.makeBreakfast())
	fmt.Println("Lunch:", m.makeLunch())
	fmt.Println("Dinner:", m.makeDinner())
}

func printFood() {
	gordon_ramsay_hells_kitchen := Restaurant{
		BestBreakfast: "steak&eggs",
		BestLunch:     "oysters",
		BestDinner:    "king crab",
	}
	cookAll(gordon_ramsay_hells_kitchen)
}
