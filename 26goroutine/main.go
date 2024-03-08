package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals []string = []string{"test"}

var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://stackoverflow.com",
		"https://amazon.com",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("All websites have been processed:", signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code in %s\n", res.StatusCode, endpoint)
	}
}
