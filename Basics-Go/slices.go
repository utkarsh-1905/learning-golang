package main

import (
	"log"
	"sort"
)

func main() {

	//Slice is like an array in Go

	//Syntax
	var mySlice []string

	//To add data in the slice at the back (push operation)

	mySlice = append(mySlice, "Utkarsh")
	mySlice = append(mySlice, "Tripathi")

	log.Println(mySlice)

	//Integer array
	var myInt []int

	myInt = append(myInt, 1)
	myInt = append(myInt, 7)
	myInt = append(myInt, 3)
	myInt = append(myInt, 2)

	log.Println(myInt)
	//sorting a slice
	sort.Ints(myInt)
	log.Println(myInt)

	//initializing the slice

	mySliceNew := []int{9, 2, 3, 4, 83, 8, 1, 26271}

	log.Println(mySliceNew)

	sort.Ints(mySliceNew)

	log.Println(mySliceNew)

	//printing a part of slice i.e. range
	log.Println(mySliceNew[3:7])

}
