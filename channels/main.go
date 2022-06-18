package main

import (
	"github.com/utkarsh-1905/learningGoChannels/randGen"
	"log"
)

func CalculateValue(intChan chan int) {
	randNumber := randGen.RandomGenerator(10000)
	//Passing a data via a channel
	intChan <- randNumber
}

//Practicing Go channels
func main() {
	//Syntax to create a channel
	intChan := make(chan int)
	//defer = Whatever comes after defer, execute it as soon as the function called here
	//in our case close(intChan)
	//Use cases =  to close files/database connections to free memory and ports
	defer close(intChan)
	//Creating a Go Routine
	//It fires off the function as a routine, passing the channel as an argument
	//The channel listens for any value to pushed in it
	//go keyword is used to start a routine
	//It executes all routines concurrently
	go CalculateValue(intChan)
	//The channel passes the value to variable whenever it hears any changes
	num := <-intChan

	log.Println(num)
}
