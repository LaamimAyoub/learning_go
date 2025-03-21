package structs

import (
	"math"
	"testing"
)

func TestPer(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("Test On Rectangle", func(t *testing.T) {
		rect := Rectangle{10.0, 10.0}
		checkArea(t, rect, 100.0)
	})

	t.Run("Test on Circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 10*10*math.Pi)
	})

}
