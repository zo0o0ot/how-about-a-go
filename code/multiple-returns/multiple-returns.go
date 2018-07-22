package main

import "fmt"

func randomishNumber() int {
	return 7 // Well, it may have been random at one point. https://xkcd.com/221/
}

func instructions() (string, string) {
	return "When I say That,", "You say Conference"
}

func callAndResponse() string {
	return "That"
}

func main() {
	foo, bar := instructions()
	fmt.Println(randomishNumber())
	fmt.Println(foo, bar)
	fmt.Println(callAndResponse())
	fmt.Println(callAndResponse())
}
