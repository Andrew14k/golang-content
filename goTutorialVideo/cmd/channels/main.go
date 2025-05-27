package main

import "fmt"

//channels hold data, are thread safe (avoid data races when reading/writing from memory), listen for data (allows block execution for certain channel events)
// allow go routines to pass information
func main() {
	var c = make(chan int) //channel with a single int value
	// c <- 1 //assigns 1 to channel
	// var i = <- c	//pops 1 from c, assigns 1 to i

	go process(c)      //make go routine
	for i := range c { //wait for something to be added to channel
		fmt.Println(i) //concurrency with process
	}

}

func process(c chan int) { //go routine to set value to channel
	for i := 0; i < 5; i++ {
		c <- i
	}
	close(c) //closes any process with channel

}
