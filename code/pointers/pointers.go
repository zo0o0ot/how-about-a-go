package main

import "fmt"

func main() {
	powerLevel := 9000
	fmt.Println(powerLevel)
	jumpStreet := &powerLevel                        // point jumpStreet to powerLevel
	fmt.Println(*jumpStreet)                         // read powerLevel through the pointer
	*jumpStreet = 21                                 // set powerLevel through the pointer
	fmt.Println(powerLevel, *jumpStreet, jumpStreet) // see the new value of variables and pointer
}
