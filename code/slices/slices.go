package main

import "fmt"

func slicesDemo() {
	ri := make([]string, 4) //Creates four strings that are empty ""
	ri[0] = "never"
	ri[3] = "you"
	ri[1] = "gonna"
	ro := make([]string, len(ri))
	copy(ro, ri) // Copies contents TO one slice FROM the other.
	ro = append(ro, "down")
	ri = append(ri, "up", "up", "down", "down")
	ri = append(ri, "left", "right", "left", "right", "b", "a", "start")
	ro[2] = "let"
	ri[2] = "give"
	fmt.Println(ri[4:])
	fmt.Println(ri[:5])
	fmt.Println(ro)
}

func main() {
	slicesDemo()
}
