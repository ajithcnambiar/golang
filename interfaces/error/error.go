package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v -- %v\n", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"Random Error",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
