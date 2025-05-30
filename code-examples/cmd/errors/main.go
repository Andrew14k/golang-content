package main

import (
	"errors"
	"fmt"
	"log"
)

func division(a, b int) (float64, error) {
	if b == 0 {
		return float64(0), errors.New("cannot divide by zero")
	}
	return float64(a) / float64(b), nil
}

func main() {
	if result, err := division(2, 5); err != nil {
		log.Println("error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

	x := 5
	defer fmt.Println("deferred:", x) // evaluated now (x = 5)
	x = 10
	fmt.Println("current:", x) // x = 10
	// output : current: 10, deferred 5
}

/*
type JsonError struct {
  StatusCode int `json:"status_code"`
  Message    any `json:"message"`
}

func (j JsonError) Error() string {
  return j.Message.(string)
}

err := fmt.Errorf("failed to open file: %v", originalErr)

return fmt.Errorf("something went wrong: %w", err) - %w wraps original error + any context you want to add to that
wrappedErr := fmt.Errorf("outer error: %w", errors.New("inner error"))
fmt.Println(errors.Unwrap(wrappedErr))
*/

/* defer, panic and recover:
defer will defer execution of something until other execution is completed
	defers are executed in a LiFo order - last defer will be first one to be executed

panic stops normal execution of a go routine, and will wind back up. other executions still go

recover allows a recovery from a panic, enabling execution to continue
*/

/*
func processTransaction(account *Account) {
    if account == nil {
        panic("account must not be nil")
    }
    // continue processing...
}
*/

/*
func safeRun() {
    defer func() {
        if r := recover(); r != nil {
            log.Println("Recovered in safeRun:", r)
        }
    }()

    mayPanic()
}

func mayPanic() {
    panic("Unexpected failure!")
}

func main() {
    fmt.Println("Starting application")
    safeRun()
    fmt.Println("Application continues running")
}
*/
