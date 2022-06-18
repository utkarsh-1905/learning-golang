package main

import (
	"errors"
	"log"
)

func main() {
	result, err := divide(5.0, 2.5)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(result)
	}
}

func divide(x float32, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("Can't divide by 0")
	}

	result = x / y
	return result, nil
}
