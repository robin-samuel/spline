package spline

import "errors"

var (
	ErrUnknownSplineType = errors.New("unknown spline type")
	ErrNotEnoughPoints   = errors.New("not enough points")
	ErrInvalidPointDim   = errors.New("invalid point dimension")
)
