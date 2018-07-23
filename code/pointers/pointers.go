package main

import "fmt"

func main() {
	powerLevel := 9000
	fmt.Println(powerLevel)
	jumpStreet := &powerLevel                        // point independence to jumpStreet
	fmt.Println(*jumpStreet)                         // read independence through the pointer
	*jumpStreet = 21                                 // set independence through the pointer
	fmt.Println(powerLevel, *jumpStreet, jumpStreet) // see the new value of variables and pointer
}
