package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")
	// how to create slices

	// var fruitSlice = []string{"Apple", "Mango", "Orange"}

	// fmt.Printf("The type of fruitSlice is %T\n", fruitSlice)

	// fruitSlice = append(fruitSlice, "Banana")
	// fmt.Println(fruitSlice)

	// // fruitSlice = append(fruitSlice[1:])
	// fmt.Println(fruitSlice)
	// fmt.Println(fruitSlice[:3])

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 345
	highScores[2] = 456
	highScores[3] = 867
	// highScores[4] = 999  	// cannot do this cuz it's out of bounds

	// fmt.Println(highScores)
	highScores = append(highScores, 999, 44)
	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)

	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	// removing value  from slice based on index

	var courses = []string{"reactJs", "js", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)
}
