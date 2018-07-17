package main

import (
	"fmt"
	"math/cmplx"
) // These are called packages.  They help you do cool stuff.

var (
	ToBe        bool       = false          // Basic types: int, string, bool
	z           complex128 = cmplx.Sqrt(-4) // Others: byte, rune, uint32, float64, uintptr
	notAssigned int                         // If not assigned, value defaults to zero.
) // You can initialize variables outside of a function

func main() {
	var ThatRhyme string = "A Poem:" // Full variable initialization
	RosesRed := true                 // Shorter syntax, also types can be inferred
	ThatVariable, GolangSweet, YouAreToo := 2, true, true
	fmt.Println(ToBe, z, notAssigned)
	fmt.Println(ThatRhyme)
	fmt.Println(RosesRed, ThatVariable, GolangSweet, YouAreToo)
}
