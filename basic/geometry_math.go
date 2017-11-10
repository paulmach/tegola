package basic

import (
	"fmt"

	"github.com/paulmach/geo"
	"github.com/paulmach/geo/project"
	"github.com/paulmach/tegola"
)

func ToWebMercator(SRID int, geometry geo.Geometry) (geo.Geometry, error) {
	switch SRID {
	default:
		return nil, fmt.Errorf("Don't know how to convert from %v to %v.", tegola.WebMercator, SRID)
	case tegola.WebMercator:
		// Instead of just returning the geometry, we are cloning it so that the user of the API can rely
		// on the result to alway be a copy. Instead of being a reference in the on instance that it's already
		// in the same SRID.
		return geo.Clone(geometry), nil
	case tegola.WGS84:
		return project.ToPlanar(geo.Clone(geometry), project.Mercator), nil
	}
}

func FromWebMercator(SRID int, geometry geo.Geometry) (geo.Geometry, error) {
	switch SRID {
	default:
		return nil, fmt.Errorf("Don't know how to convert from %v to %v.", SRID, tegola.WebMercator)
	case tegola.WebMercator:
		// Instead of just returning the geometry, we are cloning it so that the user of the API can rely
		// on the result to alway be a copy. Instead of being a reference in the on instance that it's already
		// in the same SRID.
		return geo.Clone(geometry), nil
	case tegola.WGS84:
		return project.ToLonLat(geo.Clone(geometry), project.Mercator), nil
	}
}
