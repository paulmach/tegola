package wkb

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/paulmach/geo"
)

//Polygon is a Geometry of one or more rings. The first ring is assumed to be the
// outer bounding ringer, and traversed in a clockwise manner. The remaining rings
// should be within the bounding area of the first ring, and is traversed in a counterclockwise
// manner, these represent holes in the polygon.
type Polygon struct {
	geo.Polygon
}

//Type returns the type constant for this Geometry
func (Polygon) Type() uint32 {
	return GeoPolygon
}

func (p *Polygon) Geometry() geo.Geometry {
	return p.Polygon
}

//Decode decodes the byte stream into the Geometry.
func (p *Polygon) Decode(bom binary.ByteOrder, r io.Reader) error {
	var num uint32
	if err := binary.Read(r, bom, &num); err != nil {
		return err
	}
	for i := uint32(0); i < num; i++ {
		var l = new(LineString)
		if err := l.Decode(bom, r); err != nil {
			return err
		}
		p.Polygon = append(p.Polygon, geo.Ring(l.LineString))
	}
	return nil
}

//String returns a WKT representation of the Geometry
func (p *Polygon) String() string {
	return WKT(p) // If we have a failure we don't care
}

// MultiPolygon holds multiple polygons.
type MultiPolygon struct {
	geo.MultiPolygon
}

// Type of the Geometry
func (MultiPolygon) Type() uint32 {
	return GeoMultiPolygon
}

func (mp *MultiPolygon) Geometry() geo.Geometry {
	return mp.MultiPolygon
}

// Decode decodes the binary representation of a Multipolygon and decodes it into
// a Multipolygon object.
func (mp *MultiPolygon) Decode(bom binary.ByteOrder, r io.Reader) error {
	var num uint32

	if err := binary.Read(r, bom, &num); err != nil {
		return err
	}

	for i := uint32(0); i < num; i++ {
		var p = new(Polygon)
		byteOrder, typ, err := decodeByteOrderType(r)
		if err != nil {
			return err
		}
		if typ != GeoPolygon {
			return fmt.Errorf("Expect Multipolygons to contains polygons; did not find a polygon.")
		}
		if err := p.Decode(byteOrder, r); err != nil {
			return err
		}
		mp.MultiPolygon = append(mp.MultiPolygon, p.Polygon)

	}
	return nil
}

func (mp *MultiPolygon) String() string {
	return WKT(mp.MultiPolygon) // If we have a failure we don't care
}
