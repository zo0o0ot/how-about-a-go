package main

import "fmt"

func itsNine() {
	n := 9
	if n == 9 {
		fmt.Println("Yep, it's nine.")
	} else if n == 42 {
		fmt.Println("It's the ultimate answer.")
	} else {
		fmt.Println("It's something else!")
	}

	switch n {
	case 7:
		fmt.Println("It's seven.")
	case 8:
		fmt.Println("It's eight.")
	case 9:
		fmt.Println("It's nine.")
	default:
		fmt.Println("Oh, no! 7 ate 9!")
	}
}

func main() {
	itsNine()
}
