package tests

import (
	"testing"

	"github.com/herneygr11/exercise-05/services"
)

func TestArea(t *testing.T) {
	got := 216.0
	want := services.Area(12, 18)

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}
