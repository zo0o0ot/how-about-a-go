package main

import "fmt"

// Location - similar to storing a GPS location
type Location struct {
	lat  float64
	long float64
}

// ShowTheLocation ... Shows the location.
func (l Location) ShowTheLocation() {
	fmt.Println(l)
}

func main() {
	l := Location{0.0, -89.77}
	kalahari := &l
	kalahari.lat = 43.57 // Note the use of p instead of *p
	kalahari.ShowTheLocation()
	fmt.Println(l, kalahari)
}
