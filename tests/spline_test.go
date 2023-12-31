package tests

import (
	"fmt"
	"testing"

	"github.com/robin-samuel/spline"
)

func TestLinearSpline(t *testing.T) {
	// Define points for the spline
	controllPoints := []spline.Point{
		{X: 0, Y: 2.5},
		{X: 2, Y: 4},
		{X: 3, Y: 2},
		{X: 4, Y: 1.5},
		{X: 5, Y: 6},
		{X: 6, Y: 5},
		{X: 7, Y: 3},
		{X: 9, Y: 1},
		{X: 10, Y: 2.5},
		{X: 11, Y: 7},
		{X: 9, Y: 5},
		{X: 8, Y: 6},
		{X: 7, Y: 5.5},
	}

	// Create a linear spline
	linear, err := spline.NewSpline(spline.Linear, controllPoints)
	if err != nil {
		panic(err)
	}

	// Evaluate the spline Range
	points := linear.Range(0, float64((len(controllPoints) - 1)), 0.01)
	fmt.Printf("Points on linear spline: %v\n", points)
	savePlot("linear", points, controllPoints)
}

func TestCatmullRomSpline(t *testing.T) {
	// Define points for the spline
	controllPoints := []spline.Point{
		{X: 0, Y: 2.5},
		{X: 2, Y: 4},
		{X: 3, Y: 2},
		{X: 4, Y: 1.5},
		{X: 5, Y: 6},
		{X: 6, Y: 5},
		{X: 7, Y: 3},
		{X: 9, Y: 1},
		{X: 10, Y: 2.5},
		{X: 11, Y: 7},
		{X: 9, Y: 5},
		{X: 8, Y: 6},
		{X: 7, Y: 5.5},
	}

	// Create a Catmull-Rom spline
	catmullRom, err := spline.NewSpline(spline.CatmullRom, controllPoints)
	if err != nil {
		panic(err)
	}

	// Evaluate the spline Range
	points := catmullRom.Range(0, float64((len(controllPoints) - 1)), 0.01)
	fmt.Printf("Points on Catmull-Rom spline: %v\n", points)
	savePlot("catmullrom", points, controllPoints)
}

func TestBspline(t *testing.T) {
	// Define points for the spline
	controllPoints := []spline.Point{
		{X: 0, Y: 2.5},
		{X: 2, Y: 4},
		{X: 3, Y: 2},
		{X: 4, Y: 1.5},
		{X: 5, Y: 6},
		{X: 6, Y: 5},
		{X: 7, Y: 3},
		{X: 9, Y: 1},
		{X: 10, Y: 2.5},
		{X: 11, Y: 7},
		{X: 9, Y: 5},
		{X: 8, Y: 6},
		{X: 7, Y: 5.5},
	}

	// Create a B-spline
	bspline, err := spline.NewSpline(spline.Bspline, controllPoints)
	if err != nil {
		panic(err)
	}

	// Evaluate the spline Range
	points := bspline.Range(0, float64((len(controllPoints) - 1)), 0.01)
	fmt.Printf("Points on B-spline: %v\n", points)
	savePlot("bspline", points, controllPoints)
}
