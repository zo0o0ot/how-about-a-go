package main

import (
	"fmt"
)

func showSomeMaps() {
	itsamap := make(map[string]int) // Initialize the map.
	itsamap["a"] = 1                // Steak Sauce
	itsamap["b"] = 52               // Love Shack
	itsamap["c"] = 4                // Boom
	itsamap["d"] = 20               // Natural 20! Woo!

	fmt.Println(itsamap)
	delete(itsamap, "d")         // Delete from a map by referring to its key.
	_, isPresent := itsamap["d"] // "_" refers to the first variable,  which we are going to ignore.
	// Note the two return values for the map. The second value returns if the key is present in the map.

	strangemap := make(map[rune][]string)
	strangemap[42] = []string{"The", "Answer", "To", "Life, the Universe, and Everything is", fmt.Sprintf("%c", 42)}

	fmt.Println("Roll for initiative with d20?", isPresent)
	fmt.Println(itsamap)
	fmt.Println(strangemap)

}

func main() {
	showSomeMaps()
}
