package spline

// SplineType represents the type of spline.
type SplineType int

const (
	// Linear spline type.
	Linear SplineType = iota
	// CatmullRom spline type.
	CatmullRom
	// Bspline spline type.
	Bspline
)

// Point represents a point in 2D space.
type Point struct {
	X, Y float64
}

// Spline is the interface that represents a spline.
type Spline interface {
	// Len returns the number of segments in the spline.
	Len() int
	// At returns the point at the given t value.
	At(t float64) Point
	// Range returns a slice of points from the given range.
	Range(from, to, step float64) []Point
}

// NewSpline creates a new spline based on the given type and points.
func NewSpline(t SplineType, points []Point) (Spline, error) {
	switch t {
	case Linear:
		return NewLinearSpline(points)
	case CatmullRom:
		return NewCatmullRomSpline(points, true)
	case Bspline:
		return NewBspline(points, true)
	default:
		return nil, ErrUnknownSplineType
	}
}
