package main

import "fmt"

func main() {
	notUsed := "at all" // Variable is declared but not used. Golang doesn't like that.
	fmt.Println("This won't display because of a compilation error.")
}
