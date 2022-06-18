package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Kutta",
		Breed: "Dalmation",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Jong",
		Color:         "Gray",
		NumberOfTeeth: 90,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says ", a.Says(), " and has ", a.NumberOfLegs(), " legs")
}

//The receiver type is a pointer as it makes things simpler and way faster
//And it is the best practice mentioned in the docs

func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "rwaarr"
}

func (g *Gorilla) NumberOfLegs() int {
	return 4
}
