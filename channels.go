package main

import "fmt"

func main() {
	// declare buffered channels
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// the channel can receive 2 elements and receive 2 from it
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// channel is empty, fatal error: deadlock!

	close(ch)  // close channel, if not necessary, no need to close
}