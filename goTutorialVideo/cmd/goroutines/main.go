package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutines are a way to launch multiple functions and have them execute concurrently
// concurrency - multiple tasks in process at the same time
//parallel execution - multiple CPU cores working on tasks at the same time - but tasks run in parallel
//concurrency !always= parralel

// var m = sync.Mutex{}      //mutual exclusion - controls writing to slice safely in a concurrent program
var m = sync.RWMutex{}    //prevents all other go routines from being locked, allows for read and write mutexs
var wg = sync.WaitGroup{} //counters when go routine is spawned
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    //increments counter
		go dbCall(i) //wont wait for function to complete, go allows it to move to the next step of loop
	}
	wg.Wait() //waits for counter to return to zero, meaning tasks are completed, so rest of code can execute
	fmt.Printf("\nTotal execution timeL %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	//simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log() //checks if full lock on mutex, prevent reading whilst results being written to

	///////prev code before RWMutex
	// //lock and unlock placed around results slice
	// m.Lock() //check performed to see if lock is already set, sets lock itself if it has
	// results = append(results, dbData[i])
	// m.Unlock() //now other go routines can obtain a lock if needed

	wg.Done() //decrements counter
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock() //prevents reading whilst some of the results are being
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}

//positions of locks and unlocks are very important and can affect concurrency if not positioned well
//dbCall is constant time, complex computations will have linear time if handled the above way
