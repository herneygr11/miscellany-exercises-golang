package main

import "testing"

func TestSquaring(t *testing.T) {
	got := Squaring(7)
	want := 48.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
