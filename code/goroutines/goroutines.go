package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("Boots n'")
	say("Cats n'")
}
