package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and value is %v \n", index, day)
	}

	// for _, day := range days {
	// 	fmt.Printf("index is and value is %v \n", day)
	// }

	rogueValue := 1
	// usage of break and continue statements
	for rogueValue <= 10 {
		if rogueValue == 4 {
			rogueValue++
			continue
		}
		if rogueValue == 8 {
			// break
			goto lco
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Hello World")
}
