package main

import "fmt"

func main() {

	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	// fmt.Println("List of all languages: ", languages)
	// fmt.Printf("Type of languages is %T ", languages)

	// fmt.Println("JS shorts for", languages["JS"])

	// delete(languages, "RB")
	// fmt.Println("List of all languages: ", languages)

	// looping through maps

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}
}
