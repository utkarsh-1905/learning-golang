package main

import (
	"fmt"
	"github.com/utkarsh-1905/Hello-World/pkg/handlers"
	"net/http"
)

//Port number variable
const port = ":8080"

func main() {

	//Lecture 1
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	_, err := fmt.Fprintf(w, "Hello, World!!")
	//	if err != nil {
	//		log.Println(err)
	//	}
	//})

	//Lecture 2
	//Using function handlers for requests
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/divide", handlers.Divide)
	http.HandleFunc("/hello", handlers.Hello)

	fmt.Printf("Starting at port %s", port)
	_ = http.ListenAndServe(port, nil)
}
