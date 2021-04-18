package tests

import (
	"testing"

	"github.com/herneygr11/exercise-06/services"
)

func TestArea(t *testing.T) {
	square := services.Square{
		Side: 12,
	}

	got := 144.0
	want := square.Area()

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestPerimeter(t *testing.T) {
	square := services.Square{
		Side: 12,
	}

	got := 48.0
	want := square.Perimeter()

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
