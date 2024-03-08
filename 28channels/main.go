package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// myCh <- 23
	// fmt.Println(<-myCh)
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// myCh <- 23
		val, isChannelOpen := <-myCh
		if !isChannelOpen {
			fmt.Println("You are getting flag value")
		}
		fmt.Println(val)

		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 23
		// myCh <- 24
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
