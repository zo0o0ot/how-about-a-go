package main

import "fmt"

type Location struct {
	lat  float64
	long float64
}

func main() {
	l := Location{0.0, -89.77}
	kalahari := &l
	kalahari.lat = 43.57 // Note the use of p instead of *p
	fmt.Println(l, kalahari)
}
