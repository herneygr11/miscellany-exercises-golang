package tests

import (
	"testing"

	"github.com/herneygr11/exercise-04/services"
)

func TestConverter(t *testing.T) {
	got := 11.90
	want := services.Converter(10)

	if got != want {
		t.Errorf("Error %f : %f", got, want)
	}
}
