package main

import "fmt"

func main() {
	fmt.Println("If else in golang")

	loginCount := 5
	var result string

	if loginCount > 10 {
		result = "Regular user"
	} else if loginCount <= 1 {
		result = "new user"
	} else {
		result = "unknown"
	}
	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is NOT less than 10")
	}

}
