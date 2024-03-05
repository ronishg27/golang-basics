package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointers in golang.")

	// var ptr *int;
	// fmt.Println("Value of ptr is:", ptr)

	myNumber := 23
	var ptr = &myNumber

	// fmt.Println("Value of actual pointer is:", ptr)
	// fmt.Println("Value of actual pointer is:", *ptr)

	*ptr = *ptr + 7

	fmt.Println("New value is:", myNumber)

	myNumber = 90
	println("New value is:", *ptr)

}