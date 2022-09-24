package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golaang - LearnCodeOnline.in")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)

	//receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myCh

		fmt.Println(isChannelOpen)
		fmt.Println(val)

		// fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	//send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 0
		close(myCh)
		// myCh <- 5
		// myCh <- 6
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}