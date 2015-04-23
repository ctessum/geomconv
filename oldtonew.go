package geomconv

import (
	"fmt"

	newgeom "github.com/ctessum/geom"
	oldgeom "github.com/twpayne/gogeom/geom"
)

func OldToNew(old oldgeom.T) newgeom.T {
	if old == nil {
		return nil
	}
	switch t := old.(type) {
	case oldgeom.Point:
		return newgeom.Point(old.(oldgeom.Point))
	case oldgeom.MultiPoint:
		o := old.(oldgeom.MultiPoint)
		n := make(newgeom.MultiPoint, len(o.Points))
		for i, p := range o.Points {
			n[i] = newgeom.Point(p)
		}
		return n
	case oldgeom.Polygon:
		o := old.(oldgeom.Polygon)
		n := make(newgeom.Polygon, len(o.Rings))
		for i, r := range o.Rings {
			n[i] = make([]newgeom.Point, len(r))
			for j, p := range r {
				n[i][j] = newgeom.Point(p)
			}
		}
		return n
	case oldgeom.MultiPolygon:
		o := old.(oldgeom.MultiPolygon)
		n := make(newgeom.MultiPolygon, len(o.Polygons))
		for i, p := range o.Polygons {
			n[i] = OldToNew(p).(newgeom.Polygon)
		}
		return n
	case oldgeom.LineString:
		o := old.(oldgeom.LineString)
		n := make(newgeom.LineString, len(o.Points))
		for i, p := range o.Points {
			n[i] = newgeom.Point(p)
		}
		return n
	case oldgeom.MultiLineString:
		o := old.(oldgeom.MultiLineString)
		n := make(newgeom.MultiLineString, len(o.LineStrings))
		for i, p := range o.LineStrings {
			n[i] = OldToNew(p).(newgeom.LineString)
		}
		return n
	default:
		panic(fmt.Errorf("Unsupported geom type: %v", t))
	}
}
