package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type JSONStrut struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func main() {

	myJson := `
	[
		{
			"first_name":"Utkarsh",
			"last_name":"Tripathi"
		},
		{	
			"first_name":"Elephant",
			"last_name":"Singh"
		}
	]`

	//Working with json in go is called Marshalling and Unmarshalling

	var names []JSONStrut

	//To make json from string-json file
	err := json.Unmarshal([]byte(myJson), &names)
	if err != nil {
		log.Println(err)
	}

	log.Printf("Marshalled json: %v", names)

	var myPerson []JSONStrut

	var m1 JSONStrut
	m1.FirstName = "Wally"
	m1.LastName = "West"
	myPerson = append(myPerson, m1)

	var m2 JSONStrut
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	myPerson = append(myPerson, m2)

	//Creating/Writing a JSON file from struct data in binary buffer format
	newJson, err := json.MarshalIndent(myPerson, "", "	")
	if err != nil {
		log.Println("error in marshalling", err)
	}

	//Converting buffer to string and printing
	fmt.Print(string(newJson))

}
