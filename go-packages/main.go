package main

import (
	"github.com/utkarsh-1905/myFirstPackage/helpers"
	"log"
)

func main() {
	var myVar helpers.MyStruct

	myVar.LastName = "einstein"
	myVar.FirstName = "Albert"

	log.Println(myVar.FirstName, myVar.LastName)
}
