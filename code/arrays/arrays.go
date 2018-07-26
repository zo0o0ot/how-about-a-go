package main

import "fmt"

func arraysDemo() {
	var a [9]int                           // This is an array of length 9.  It is filled with zeros
	b := [9]int{1, 2, 4, 3, 0, 9, 8, 6, 7} // Arrays are fixed.  They can never be expanded.
	b[2] = 5                               // Arrays are zero-indexed.
	fmt.Printf("%T \n", a)                 // Show the type of variable.
	fmt.Println(a)                         // Print the contents of the array.
	// Array ranges can be referenced, [2:6] refers to array positions 2-5.
	// Either position in a range can be omitted.
	fmt.Println(b[6:], "-", b[2:6])

	for x, y := range b { // Range does something similar to foreach.
		fmt.Printf("%d-%d ", x, y)
	}
	fmt.Println("")
}

func main() {
	arraysDemo()
}
