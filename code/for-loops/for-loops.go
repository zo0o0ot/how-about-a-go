package main

import "fmt"

func theCount() {
	for i := 1; i < 6; i++ {
		fmt.Println(i, "!!!!! AH! AH! AH!")
	}
	n := 6
	for ; n <= 10; n++ {
		fmt.Println(n, "!!!!! AH! AH! AH!")
	}
	fmt.Println("NOW! COUNT BY TWOS!")
	for ; ; n++ {
		if n%2 == 0 {
			fmt.Println(n, "!!!!! AH! AH! AH!")
		} else if n < 20 {
			continue
		} else {
			return
		}
	}
}

func main() {
	fmt.Println("I am the Count, and I love to count!")
	theCount()
}
