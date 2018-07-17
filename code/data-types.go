package main

import (
	"fmt"
	"math/cmplx"
) // These are called packages.  They help you do cool stuff.

var (
	ToBe        bool       = false
	z           complex128 = cmplx.Sqrt(-4) // Fancy types exist.
	notAssigned int                         // If not assigned, value defaults to zero.
) // You can initialize variables outside of a function

func main() {
	var ThatString string = "That" // Full variable initialization
	d6 := 6                        // Shorter syntax, also types can be inferred
	RosesRed, VioletsBlue, GolangSweet, YouAreToo := true, true, true, true
	fmt.Println(ToBe, z, notAssigned)
	fmt.Println(ThatString, d6)
	fmt.Println(RosesRed, VioletsBlue, GolangSweet, YouAreToo)
}
