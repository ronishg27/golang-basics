package main

import "fmt"

func main() {
	//arrays in golang

	// var fruitList [4]string //declaring an array

	// fruitList[0] = "Apple"
	// fruitList[1] = "Orange"
	// fruitList[3] = "Banana"

 var fruitList2 = [4]string{"Apple", "Orange", "Banana"}

	fmt.Println("Fruit list is:", fruitList2)
	fmt.Printf("Type of fruit list is: %T\n", fruitList2)

}