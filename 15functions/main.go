package main

import "fmt"

func main() {

	fmt.Println("Functions in Golang")

	sum, _ := add(10, 20, 40)
	fmt.Printf("Sum of all numbers is: %d", sum)
}

func add(values ...int) (int, string) {

	total := 0
	for _, value := range values {
		total += value
	}

	return total, "Added successfully"
}
