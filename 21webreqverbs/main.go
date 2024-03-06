package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("web request verbs in golang")

	// performGETRequest()
	// performPOSTJsonRequest()
	performPOSTFormRequest()
}

func performGETRequest() {
	const myUrl = "http://localhost:8000/get"

	res, err := http.Get(myUrl)
	checkNilError(err)
	defer res.Body.Close()

	fmt.Println("Status code: ", res.StatusCode)
	fmt.Println("Content length is: ", res.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(res.Body)

	byteCount, _ := responseString.Write(content)
	fmt.Println(byteCount)

	fmt.Println(responseString.String())
	// fmt.Println("Content is: ", string(content))

}

func performPOSTJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload

	reqBody := strings.NewReader(`
		{
			"foo1": "bar1",
			"foo2": "bar2",
			"foo3": "bar3"
		}
		`)

	res, err := http.Post(myUrl, "application/json", reqBody)
	checkNilError(err)
	defer res.Body.Close()

	dataBytes, _ := io.ReadAll(res.Body)
	checkNilError(err)

	fmt.Println(string(dataBytes))
}

func performPOSTFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// form data

	data := url.Values{}
	data.Add("firstName", "Ronish")
	data.Add("lastName", "Ghim")
	data.Add("email", "ronish@go.dev")

	res, err := http.PostForm(myUrl, data)
	checkNilError(err)
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	fmt.Println("Content is: ", string(content))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
