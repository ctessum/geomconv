package geom2rtree

// Package geom2rtree calculates the rtreego bounding box of a Geom shape.

import (
	"github.com/patrick-higgins/rtreego"
	"github.com/twpayne/gogeom/geom"
)

// Calculate the rtreego bounding box of a geom shape.
func GeomToRect(g geom.T) (*rtreego.Rect, error) {
	b := geom.NewBounds()
	b = g.Bounds(b)
	p := rtreego.Point{b.Min.X, b.Min.Y}
	lengths := [3]float64{b.Max.X - b.Min.X,
		b.Max.Y - b.Min.Y}
	if lengths[0] == 0. {
		lengths[0] = 1.e-6
	}
	if lengths[1] == 0. {
		lengths[1] = 1.e-6
	}
	lengths[2] = 1.e-6
	bounds, err := rtreego.NewRect(p, lengths)
	return &bounds, err
}
