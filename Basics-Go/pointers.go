package pointers

import (
	"fmt"
)

func main() {
	var myString string

	myString = "hello"

	fmt.Println(myString)
	changeUsingPointer(&myString)
	fmt.Println(myString)
}

func changeUsingPointer(s *string) {
	newString := "Hola, Mundo!"
	*s = newString
}
