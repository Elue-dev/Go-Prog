package main

import (
	"fmt"

	"github.com.elue-dev/go-program/helpers"
)

type Animal interface {
	Says() 		        string
	NumberofLegs() 		int
}

type Dog struct {
	Name 		string
	Breed 		string
}

type Gorilla struct {
	Name 			string
	Color 			string
	NumberOfTeeth 	int
}

func main() {
	dog := Dog{
		Name: "Swanson",
		Breed: "Boerbel",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Swanson",
		Color: "Boerbel",
		NumberOfTeeth: 3,
	}

	PrintInfo(&gorilla)

	var myVar helpers.SomeType
	myVar.Name = "Wisdom"
	fmt.Println((myVar))
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberofLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}

func (d *Dog) NumberofLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberofLegs() int {
	return 2
}