package main

import (
	"fmt"
)

func showSomeMaps() {
	itsamap := make(map[string]int)
	itsamap["a"] = 1
	itsamap["b"] = 52
	itsamap["c"] = 4
	itsamap["d"] = 20

	fmt.Println(itsamap)
	delete(itsamap, "d")         //
	_, isPresent := itsamap["d"] // Note the two return values.
	// Also "_" refers to a variable that we are going to ignore.

	strangemap := make(map[rune]string)

	fmt.Println("Roll for initiative with d20?", isPresent)
	fmt.Println(itsamap)
	fmt.Println(strangemap)

}

func main() {
	showSomeMaps()
}
