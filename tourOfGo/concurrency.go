// CONCURRENCY

// go routine is a lightweight thread managed by the go runtime
// go routines run in the same address space, so access to shared memory must be synchronised
// sync package provides useful primitives
// package main
// import (
// 	"fmt"
// 	"time"
// )
// func say(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(100 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }
// func main() {
// 	go say("world")
// 	say("hello")
// }

// channels - great for communication among go routines
//typed conduit in which you can send and receive values with the channel operator <-
//code below syms the numbers in a slice, distributing work between go routines, once both completed, calcs final result
// func sum(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 	}
// 	c <- sum // send sum to c
// }

// func main() {
// 	s := []int{7, 2, 8, -9, 4, 0}
// 	c := make(chan int)
// 	go sum(s[:len(s)/2], c)
// 	go sum(s[len(s)/2:], c)
// 	x, y := <-c, <-c // receive from c
// 	fmt.Println(x, y, x+y)
// }

// buffered channels
// ch := make(chan int, 100) // sends to buffered channel only block when buffer is full

// closing channels are always necessary, only when the receiver must be told there are no more values coming, such as terminating a range loop
// func fibonacci(n int, c chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < n; i++ {
// 		c <- x
// 		x, y = y, x+y
// 	}
// 	close(c)
// }

// func main() {
// 	c := make(chan int, 10)
// 	go fibonacci(cap(c), c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// select statement - lets a go routine wait on multiple communication operations
// select blocks until one of its cases can run, then it executes that case. chooses one at random if multiple are ready
// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	quit := make(chan int)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println(<-c)
// 		}
// 		quit <- 0
// 	}()
// 	fibonacci(c, quit)
// }

// default case in select is run if no other case is ready
// select {
// case i := <-c:
//     // use i
// default:
//     // receiving from c would block
// }

// sync.Mutex - used when go routines dont need to communicate
// use mutex syncs if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts - mutual exclusion
// can define a block of code to be executed in mutual exclusion by surrounding it was a call to Lock and Unlock as show below.
// must use defer to ensure the mutex will be unlocked as in the Value method

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // SafeCounter is safe to use concurrently.
// type SafeCounter struct {
// 	mu sync.Mutex
// 	v  map[string]int
// }

// // Inc increments the counter for the given key.
// func (c *SafeCounter) Inc(key string) {
// 	c.mu.Lock()
// 	// Lock so only one goroutine at a time can access the map c.v.
// 	c.v[key]++
// 	c.mu.Unlock()
// }

// // Value returns the current value of the counter for the given key.
// func (c *SafeCounter) Value(key string) int {
// 	c.mu.Lock()
// 	// Lock so only one goroutine at a time can access the map c.v.
// 	defer c.mu.Unlock()
// 	return c.v[key]
// }

// func main() {
// 	c := SafeCounter{v: make(map[string]int)}
// 	for i := 0; i < 1000; i++ {
// 		go c.Inc("somekey")
// 	}

// 	time.Sleep(time.Second)
// 	fmt.Println(c.Value("somekey"))
// }
