package main

import "testing"

func TestSum(t *testing.T) {
	got := Sum(10, 5)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
