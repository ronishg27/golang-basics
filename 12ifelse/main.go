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

}
