package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race condition")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}

	// wg.Add(3)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("1st routine")
		score = append(score, 1)
		mut.Unlock()
		// wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("2nd routine")
		score = append(score, 2)
		// wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		mut.Lock()
		fmt.Println("3rd routine")
		score = append(score, 3)
		mut.Unlock()
		// wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)

}
