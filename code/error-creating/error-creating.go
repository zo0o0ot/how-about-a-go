package main

import (
	"errors"
	"fmt"
)

func alwaysErrors() (string, error) {
	return "I AM ERROR", errors.New("Error: Conference Foul- Not enough memes")
}

func main() {
	_, err := alwaysErrors()
	if err != nil {
		fmt.Println(err)
	}

}
