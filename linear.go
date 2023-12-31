package spline

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// linearSpline represents a linear spline.
type linearSpline struct {
	points []*mat.VecDense
}

// NewLinearSpline creates a new linear spline based on the given points.
func NewLinearSpline(points []Point) (Spline, error) {
	if len(points) < 2 {
		return nil, ErrNotEnoughPoints
	}

	// convert points to mat.VecDense
	matPoints := make([]*mat.VecDense, len(points))
	for i, point := range points {
		matPoints[i] = mat.NewVecDense(2, []float64{point.X, point.Y})
	}

	return &linearSpline{matPoints}, nil
}

func (s *linearSpline) Len() int {
	return len(s.points)
}

func (s *linearSpline) At(t float64) Point {
	numSegments := len(s.points) - 1

	if t > float64(numSegments) {
		t = float64(numSegments)
	} else if t < 0 {
		t = 0
	}

	segmentIndex := int(math.Floor(t))
	if segmentIndex >= numSegments {
		segmentIndex = numSegments - 1
	}

	localT := t - float64(segmentIndex)

	startPoint := s.points[segmentIndex]
	endPoint := s.points[segmentIndex+1]

	interpPoint := mat.NewVecDense(2, nil)
	interpPoint.ScaleVec(1-localT, startPoint)
	tempVec := mat.NewVecDense(2, nil)
	tempVec.ScaleVec(localT, endPoint)
	interpPoint.AddVec(interpPoint, tempVec)

	return Point{math.Round(interpPoint.AtVec(0)*100000) / 100000, math.Round(interpPoint.AtVec(1)*100000) / 100000}
}

func (s *linearSpline) Range(start, end, step float64) []Point {
	if start > end {
		start, end = end, start
	}
	n := int((end-start)/step) + 1
	t := start
	var points []Point
	for i := 0; i < n; i++ {
		points = append(points, s.At(t))
		t += step
	}
	return points
}
