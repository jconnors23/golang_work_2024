// constructors are methods used to initialize object properties and are a common feature of object oriented programming
// Go does not have defined contstructor methods, but similar functionality can be acheived through custom functions

package main

import (
	"fmt"
)

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) NewDog(name string, breed string) {
	d.Name = name
	d.Breed = breed
}

func define_dogs() {
	dog := new(Dog)
	dog.NewDog("Max", "poodle")
	fmt.Printf("%s\n", dog.Name)
	fmt.Printf("%s\n", dog.Breed)
}
