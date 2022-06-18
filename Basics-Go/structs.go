package structs

import (
	"log"
	"time"
)

//Declaring structs in GO

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	birthDay    time.Time
}

func main() {

	user := User{
		FirstName:   "Utkarsh",
		LastName:    "Tripathi",
		PhoneNumber: "123",
		Age:         19,
	}

	log.Println(user.FirstName)
}
