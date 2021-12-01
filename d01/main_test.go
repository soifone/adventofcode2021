package main

import (
	"testing"
)

func TestSlidingWindow(t *testing.T) {
	input := []int{
		199, 200, 208, 210, 200, 207, 240, 269, 260, 263,
	}

	want := []int{
		607, 618, 618, 617, 647, 716, 769, 792,
	}
	got := slidingWindow(input)

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Fatalf("Got %d\nWant %d", got[i], want[i])
		}

	}
}

func TestCountOnIncrease(t *testing.T) {
	input := []int{
		607, 618, 618, 617, 647, 716, 769, 792,
	}
	got := countOnIncrease(input)
	if got != 5 {
		t.Fatalf("Got %d\nWant %d", got, 5)
	}
}
