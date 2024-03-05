package main

import "fmt"

func main() {
	// getting input from user
	var x, y int
	fmt.Println("Enter two numbers: ")
	fmt.Scanf("%d %d", &x, &y)
	fmt.Printf("Sum of %d and %d is %d", x, y, add(x, y))

}

func add(x int, y int) int {
	return x + y
}