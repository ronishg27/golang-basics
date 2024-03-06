package main

import (
	"fmt"
	"net/url"
)

// const myUrl string = "https://github.com:3000/trending?q=golang"
const myUrl string = "https://wikipedia.org/w/index.php?title=Main_Page&action=edit"

func main() {
	fmt.Println("URL handling in golang")

	// parsing
	result, err := url.Parse(myUrl)
	checkNilError(err)
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	// fmt.Println(result.RawQuery)

	qParams := result.Query()
	fmt.Printf("The type of qParams is %T\n", qParams)
	// fmt.Println("query Params are: ", qParams["action"])

	for key, val := range qParams {
		// fmt.Println("query params are: ", val)
		fmt.Printf("key is %v and value is %v \n", (key), val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "wikipedia.org",
		Path:     "/w/index.php",
		RawQuery: "title=Main_Page",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
