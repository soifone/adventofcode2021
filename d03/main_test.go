package main

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {

	input := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	o, c := oxygenAndScrubber(input)
	if o != 23 || c != 10 {
		t.Fatalf("Oxygen\nWant: %d\nGot: %d\n\no2 Scrubber\nWant: %d\nGot: %d", 23, o, 10, c)
	}
}
