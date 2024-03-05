package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	currentTime := time.Now()

	// fmt.Println(currentTime)

	fmt.Println(currentTime.Format("01-02-2006 03:04:05 Monday"))
	// fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))


	// createdDate := time.Date(2003, time.July, 27, 17, 15, 0, 0, time.UTC)
	// fmt.Println(createdDate)
	// fmt.Println(createdDate.Format("01-02-2006 03:04:05 Monday"))


}