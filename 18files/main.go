package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in file"
	file, error := os.Create("./myFile.txt")

	// if error != nil {
	// 	panic(error)
	// }
	checkNilError(error)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myFile.txt")
}

func readFile(filename string) {
	// ioutil.ReadFile(filename)  //!deprecated

	dataByte, err := os.ReadFile(filename)

	checkNilError(err)
	fmt.Println("Text data inside the file is \n", string(dataByte))
}

func checkNilError(e error) {
	if e != nil {
		panic(e)
	}
}
