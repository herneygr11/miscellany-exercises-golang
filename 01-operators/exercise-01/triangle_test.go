package main

import "testing"

func TestTriangle(t *testing.T) {
	triangle := Triangle{base: 12, height: 15}

	got := triangle.Area()
	want := 90.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
