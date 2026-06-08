package main

import (
	"fmt"
	"time"
)

// type error interface{ //already implemented in FMT package
//  Error() string
// }

type myError struct {
	when time.Time
	what string
}

func (e *myError) Error() string { // this method implements error interface
	return fmt.Sprintf("at %v %s ", e.when, e.what)
}

func run() error {
	return &myError{
		time.Now(),
		"some error occured",
	}
}

type fileError struct {
	when time.Time
	what string
}

func (e *fileError) Error() string { // this method implements error interface
	return fmt.Sprintf("at %v %s ", e.when, e.what)
}

func run2() error {
	return &fileError{
		time.Now(),
		"some error occured",
	}
}

func main() {
	fmt.Println("error interface")

	if err := run(); err != nil {
		fmt.Println(err)
	}
	if err := run2(); err != nil {
		fmt.Println(err)
	}

	excersice()
}
