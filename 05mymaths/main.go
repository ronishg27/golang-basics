package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// fmt.Println("Welcome to maths in golang")

	// var num1 int = 10
	// var num2 float64 = 20

	// sum:=num1+num2

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// crypto rand

	randomValue, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomValue)

}