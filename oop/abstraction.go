// abstraction is about hiding implementation details of code and removing unnecessary object details
// interfaces and structs in Go provide a means to achievieng this
package main

import "fmt"

type Food interface {
	GetCalories() int
	GetBrand() string
}

type Chips struct {
	Calories int
	Brand    string
}

type Poultry struct {
	Calories int
	Brand    string
}

func (c Chips) GetCalories() int {
	return c.Calories
}

func (c Chips) GetBrand() string {
	return c.Brand
}

func (p Poultry) GetCalories() int {
	return p.Calories
}

func (p Poultry) GetBrand() string {
	return p.Brand
}

// main uses food interface without knowing type of food = abstraction(hiding details of food types)
func main() {
	f := Chips{Calories: 200, Brand: "Lays"}
	fmt.Println(f.GetCalories())
	fmt.Println(f.GetBrand())
}
