package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"courseName"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Introducing JSON in golang")
	// EncodeJson()
	DecodeJson()
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodeJson() {
	WebCourse := []course{
		{
			"ReactJS", 299, "Udemy", "abc123", []string{"web", "react"},
		},
		{
			"AngularJS", 399, "Udemy", "xyz123", []string{"full-stack", "angular"},
		},
		{
			"VueJS", 499, "Udemy", "pqr123", nil,
		},
	}

	// package this data as JSON data

	finalJson, err := json.MarshalIndent(WebCourse, "", "    ")
	checkNilError(err)
	fmt.Printf("%s\n", string(finalJson))
}

// consuming JSON
func DecodeJson() {
	jsonData := []byte(`
		{
			"courseName": "ReactJS",
			"Price": 299,
			"website": "Udemy",
			"tags": ["web", "react"]
		}
	`)

	/*
		var jsCourse course
			isValidJson := json.Valid(jsonData)
			if isValidJson {
				fmt.Println("JSON is valid")
				json.Unmarshal(jsonData, &jsCourse)
				fmt.Printf("%#v\n", jsCourse)
			} else {
				fmt.Println("JSON IS NOT VALID")
			}
	*/
	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	// fmt.Printf("%#v\n", myData)
	// fmt.Printf("%+v\n", myData)

	for k, v := range myData {
		fmt.Printf("Key: %v\tValue: %v\tType: %T\n", k, v, v)
	}

}
