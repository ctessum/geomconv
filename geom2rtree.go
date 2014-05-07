package geom2rtree

// Package geom2rtree calculates the rtreego bounding box of a Geom shape.

import (
	"github.com/dhconnelly/rtreego"
	"github.com/twpayne/gogeom/geom"
)

// Calculate the rtreego bounding box of a geom shape.
func GeomToRect(g geom.T) (bounds *rtreego.Rect, err error) {
	b := geom.NewBounds()
	b = g.Bounds(b)
	p := rtreego.Point{b.Min.X, b.Min.Y}
	lengths := []float64{b.Max.X - b.Min.X,
		b.Max.Y - b.Min.Y}
	if lengths[0] == 0. {
		lengths[0] = 1.e-6
	}
	if lengths[1] == 0. {
		lengths[1] = 1.e-6
	}
	bounds, err = rtreego.NewRect(p, lengths)
	if err != nil {
		return
	}
	return
}
