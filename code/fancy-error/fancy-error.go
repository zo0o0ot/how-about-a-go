package main

import (
	"fmt"
	"time"
)

// START OMIT

// MyError : It's fancy
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("At %02d:%02d on %d-%02d-%02d %s",
		e.When.Hour(), e.When.Minute(), e.When.Year(), e.When.Month(), e.When.Day(), e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		",\nYou tried to run code during a presentation.\nNot a good idea.",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// END OMIT
