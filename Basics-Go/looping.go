package main

import "log"

func main() {
	//Simple for loop
	for i := 0; i < 10; i++ {
		log.Println(i)
	}

	//Ranging

	names := []string{"Utkarsh", "Tripathi", "Kumar"}
	for i, name := range names {
		log.Println("Index:", i, "Name:", name)
	}

	//Use _ do ignore return values

	for _, name := range names {
		log.Println("Name:", name)
	}

	type nUser struct {
		FirstName   string
		LastName    string
		PhoneNumber string
	}

	var users []nUser //Slice of structs
	users = append(users, nUser{"Buffalo", "Gupta", "123"})
	users = append(users, nUser{"Elephant", "Sharma", "456"})
	users = append(users, nUser{"Horse", "Singh", "789"})
	users = append(users, nUser{"Monkey", "Kumar", "012"})

	for _, user := range users {
		log.Println(user)
	}
}
