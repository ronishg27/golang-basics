package main

import (
	"fmt"
	"io"
	"net/http"
)

// const url = "https://lco.dev"

const url = "https://guthib.com"

func main() {
	fmt.Println("Web requests in golang")

	response, err := http.Get(url)
	checkNilError(err)
	fmt.Printf("type of response is: %T\n", response)
	// fmt.Println(response)

	defer response.Body.Close() // closing the connection

	dataBytes, err := io.ReadAll(response.Body)
	checkNilError(err)
	fmt.Println("Data received is: ", string(dataBytes))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
