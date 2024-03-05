package main

import "fmt"

func main() {
	var username string = "johndoe"
	fmt.Println("Hello, " + username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isAdult bool = true
	fmt.Println(isAdult)
	fmt.Printf("Variable is of type: %T \n", isAdult)


	// Implicit type
	var age = 30
	fmt.Println(age)
	fmt.Printf("Variable is of type: %T \n", age)

	// No var style
	lastName := "smith"
	fmt.Println(lastName)
	fmt.Printf("Variable is of type: %T \n", lastName)

	// := is only used inside functions

}