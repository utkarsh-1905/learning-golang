package struct_functions

import "fmt"

type MyUser struct {
	FirstName string
}

//Associating functions with structs
//We pass a pointer object of struct which binds to the function

func (m *MyUser) printFirstName() {
	fmt.Println(m.FirstName)
}

func main() {

	var user1 MyUser

	user1.FirstName = "Utkarsh"

	user1.printFirstName()

}
