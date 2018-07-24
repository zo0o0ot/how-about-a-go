package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("OK")
	err := errors.New("Conference Foul: Too many words")
	if err != nil {
		fmt.Println(err)
	}

}
