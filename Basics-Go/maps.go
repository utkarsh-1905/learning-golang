package maps

import (
	"log"
)

type newUser struct {
	FirstName string
}

func main() {

	//Creating a Map in Go
	//Syntax
	myMap := make(map[string]string)
	//make is built-in function
	//[string] => key
	//string => value

	myMap["hello"] = "world"

	log.Println(myMap["hello"])

	//over-writing is allowed

	myMap["hello"] = "not world"

	log.Println(myMap["hello"])

	//Map for user-defined data type

	userMap := make(map[int]newUser)

	me := newUser{
		FirstName: "Utkarsh",
	}

	userMap[1] = me

	log.Println(userMap[1].FirstName)

	//We can pass the map itself as it is immutable and constant

	//To store any type of data in map

	//myMap3 := make(map[int]interface{})

	//Problem = We will need to cast the data everytime we pull it from the map
}
