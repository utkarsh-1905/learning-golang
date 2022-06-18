package variable_functions

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!!")
	//Variable declaration
	//we need to specify type of variable
	var whatToSay string
	var i int

	i = 10
	whatToSay = "Goodbye, cruel world!!"

	fmt.Println(whatToSay, i)

	//dynamically assigning type to variable when return type is unknown
	// := is used to assign some value to variable when typeof(value) is unknown
	//more or less same as destructuring in js

	whatWasSaid, someNumber := saySomething()

	fmt.Println(whatWasSaid, someNumber)
}

//function declaration
//return type is string,int

func saySomething() (string, int) {
	return "a", 30
}
