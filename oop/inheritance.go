// inheritance is the process of allowing the properties of
// one class be inheritaed by another class, creating a relatinoship between these classes and allowing for code reusabiltiy
// structs can be used in go to acheive similar functinoality to classes
// the shirt being initialized here in main has direct access and 'inherits' the fields of the Shirt struct

package main

import (
	"fmt"
)

type Clothing struct {
	Color string
	Size  string
}

type Shirt struct {
	Clothing
}

func PrintShirt() {
	shirt := &Shirt{
		Clothing{
			Color: "blue",
			Size:  "large",
		},
	}
	fmt.Println(shirt.Color)
	fmt.Println(shirt.Size)

}
