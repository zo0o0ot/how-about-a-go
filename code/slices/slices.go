package main

import "fmt"

func slicesDemo() {
	ri := make([]string, 4)
	ri[0] = "never"
	ri[3] = "you"
	ri[1] = "gonna"
	ro := make([]string, len(ri))
	copy(ro, ri)
	ro = append(ro, "down")
	ri = append(ri, "up")
	ro[2] = "let"
	ri[2] = "give"
	fmt.Println(ri)
	fmt.Println(ro)
}

func main() {
	slicesDemo()
}
