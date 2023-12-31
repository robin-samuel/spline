package spline

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// catmullRomSpline represents a Catmull-Rom spline.
type catmullRomSpline struct {
	points []*mat.VecDense
}

// NewCatmullRomSpline creates a new Catmull-Rom spline based on the given points.
func NewCatmullRomSpline(points []Point, ghost bool) (Spline, error) {
	if len(points) < 2 {
		return nil, ErrNotEnoughPoints
	}

	// convert points to mat.VecDense
	matPoints := make([]*mat.VecDense, len(points))
	for i, point := range points {
		matPoints[i] = mat.NewVecDense(2, []float64{point.X, point.Y})
	}

	if ghost {
		// add mirrored points to the beginning and end of the list
		ghostPoint0 := mat.NewVecDense(2, nil)
		ghostPoint0.SubVec(matPoints[0], matPoints[1])
		ghostPoint0.ScaleVec(-1, ghostPoint0)
		ghostPoint0.AddVec(matPoints[0], ghostPoint0)
		matPoints = append([]*mat.VecDense{ghostPoint0}, matPoints...)

		ghostPoint1 := mat.NewVecDense(2, nil)
		ghostPoint1.SubVec(matPoints[len(matPoints)-1], matPoints[len(matPoints)-2])
		ghostPoint1.ScaleVec(-1, ghostPoint1)
		ghostPoint1.AddVec(matPoints[len(matPoints)-1], ghostPoint1)
		matPoints = append(matPoints, ghostPoint1)
	}

	return &catmullRomSpline{matPoints}, nil
}

func (s *catmullRomSpline) Len() int {
	return len(s.points) - 2
}

func (s *catmullRomSpline) At(t float64) Point {
	numSegments := len(s.points) - 2

	t = math.Max(0, math.Min(t, float64(numSegments-1)))

	if t > float64(numSegments) {
		t = float64(numSegments)
	} else if t < 0 {
		t = 0
	}

	segmentIndex := int(math.Floor(t))
	if segmentIndex > numSegments-2 {
		segmentIndex = numSegments - 2
	}
	localT := t - float64(segmentIndex)

	// Basis matrix for Catmull-Rom spline
	basisMatrix := mat.NewDense(4, 4, []float64{
		0, 2, 0, 0,
		-1, 0, 1, 0,
		2, -5, 4, -1,
		-1, 3, -3, 1,
	})
	basisMatrix.Scale(0.5, basisMatrix)

	tVec := mat.NewDense(1, 4, []float64{1, localT, localT * localT, localT * localT * localT})

	gVec := mat.NewDense(4, 2, []float64{
		s.points[segmentIndex].AtVec(0), s.points[segmentIndex].AtVec(1),
		s.points[segmentIndex+1].AtVec(0), s.points[segmentIndex+1].AtVec(1),
		s.points[segmentIndex+2].AtVec(0), s.points[segmentIndex+2].AtVec(1),
		s.points[segmentIndex+3].AtVec(0), s.points[segmentIndex+3].AtVec(1),
	})

	intermediate := mat.NewDense(1, 4, nil)
	intermediate.Mul(tVec, basisMatrix)

	result := mat.NewDense(1, 2, nil)
	result.Mul(intermediate, gVec)

	return Point{math.Round(result.At(0, 0)*100000) / 100000, math.Round(result.At(0, 1)*100000) / 100000}
}

func (s *catmullRomSpline) Range(start, end, step float64) []Point {
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
