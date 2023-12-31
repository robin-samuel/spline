package tests

import (
	"github.com/robin-samuel/spline"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type xyPoints []spline.Point

func (points xyPoints) Len() int {
	return len(points)
}

func (points xyPoints) XY(i int) (float64, float64) {
	return points[i].X, points[i].Y
}

func savePlot(name string, points, controlPoints xyPoints) {
	p := plot.New()

	p.Title.Text = name
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	s1, err := plotter.NewLine(points)
	if err != nil {
		panic(err)
	}
	s1.Color = plotutil.Color(1)

	s2, err := plotter.NewScatter(controlPoints)
	if err != nil {
		panic(err)
	}
	s2.Color = plotutil.Color(0)
	s2.Shape = draw.CircleGlyph{}

	p.Add(s1, s2)

	if err := p.Save(16*vg.Centimeter, 12*vg.Centimeter, name+".png"); err != nil {
		panic(err)
	}
}
