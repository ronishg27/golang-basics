package main

import "fmt"

func main() {
	defer fmt.Println("in golang")
	fmt.Println("Defer")
	defer fmt.Println("one")
	defer fmt.Println("two")
	// defer, two, one, in golang
	myDefer()
}

// defer order = last in first out

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// can be used for freeing memory at the top
