package main

import (
	"math"
	"testing"
)

func Test_CalculatePerimeter(t *testing.T) {
	got := Perimeter(Rectangle{width: 5.0, height: 5.0})
	want := 20.0

	if got != want {
		t.Errorf("got %.2f; want %.2f", got, want)
	}
}

func Test_Area(t *testing.T) {
	t.Run("rectangle", func(t *testing.T) {
		shape := Rectangle{width: 5.0, height: 5.0}
		got := shape.Area()
		want := 25.0

		if got != want {
			t.Errorf("%#v got %.2f; want %.2f", shape, got, want)
		}
	})

	t.Run("circle", func(t *testing.T) {
		shape := Circle{radius: 2.5}
		got := shape.Area()
		want := 19.63

		if !assertFloatsAreEqual(t, toFixed(got, 2), toFixed(want, 2)) {
			t.Errorf("%#v got %.2f; want %.2f", shape, got, want)
		}
	})

	t.Run("triangle", func(t *testing.T) {
		shape := Triangle{base: 5, height: 4}
		got := shape.Area()
		want := 10.0

		if !assertFloatsAreEqual(t, toFixed(got, 2), toFixed(want, 2)) {
			t.Errorf("%#v got %.2f; want %.2f", shape, got, want)
		}
	})
}

const epsilon = 1e-12

func assertFloatsAreEqual(t *testing.T, a float64, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
