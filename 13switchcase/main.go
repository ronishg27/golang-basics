package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("Switch case in golang")

	// rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	// fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1, you can open.")
	case 2:
		fmt.Println("You can move 2 spots.")
	case 3:
		fmt.Println("You can move 3 spots.")
	case 4:
		fmt.Println("You can move 4 spots.")
	case 5:
		fmt.Println("You can move 5 spots.")
		fallthrough
	case 6:
		fmt.Println("You can move 6 spots and roll the dice again.")
		fallthrough
	default:
		fmt.Println("What was that?")
	}
}
