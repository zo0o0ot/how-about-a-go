package main

import "testing"

func TestWoo(t *testing.T) {
	foo := Woo()
	if foo != "woo!" {
		t.Errorf("Woo was incorrect, got: %s, want: %s.", foo, "woo!")
	}
}
