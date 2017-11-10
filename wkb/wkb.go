//Package wkb is for decoding ESRI's Well Known Binary (WKB) format for OGC geometry (WKBGeometry)
//	sepcification at http://edndoc.esri.com/arcsde/9.1/general_topics/wkb_representation.htm
// There are a few types supported by the specification. Each general type is in it's own file.
// So, to find the implementation of Point (and MultiPoint) it will be located in the point.go
// file. Each of the basic type here adhere to the tegola.Geometry interface. So, a wkb point
// is, also, a tegola.Point
package wkb

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/paulmach/geo"
)

//  geometry types
// http://edndoc.esri.com/arcsde/9.1/general_topics/wkb_representation.htm
const (
	GeoPoint              uint32 = 1
	GeoLineString                = 2
	GeoPolygon                   = 3
	GeoMultiPoint                = 4
	GeoMultiLineString           = 5
	GeoMultiPolygon              = 6
	GeoGeometryCollection        = 7
)

// Geometry describes a basic Geometry type that can decode it's self.
type Geometry interface {
	Decode(bom binary.ByteOrder, r io.Reader) error
	Geometry() geo.Geometry
	Type() uint32
}

func decodeByteOrderType(r io.Reader) (byteOrder binary.ByteOrder, typ uint32, err error) {
	var bom = make([]byte, 1, 1)
	// the bom is the first byte
	if _, err = r.Read(bom); err != nil {
		return byteOrder, typ, err
	}

	if bom[0] == 0 {
		byteOrder = binary.BigEndian
	} else {
		byteOrder = binary.LittleEndian
	}

	// Reading the type which is 4 bytes
	err = binary.Read(r, byteOrder, &typ)
	return byteOrder, typ, err
}

func encode(bom binary.ByteOrder, geometry geo.Geometry) (data []interface{}) {

	if bom == binary.LittleEndian {
		data = append(data, byte(1))
	} else {
		data = append(data, byte(0))
	}
	switch g := geometry.(type) {
	default:
		return nil
	case geo.Point:
		data = append(data, GeoPoint)
		data = append(data, g.X(), g.Y())
		return data
	case geo.MultiPoint:
		data = append(data, GeoMultiPoint)
		if len(g) == 0 {
			return data
		}
		for _, p := range g {
			data = append(data, encode(bom, p)...)
		}
		return data
	case geo.LineString:
		data = append(data, GeoLineString)
		data = append(data, uint32(len(g))) // Number of points in the line string
		for i := range g {
			data = append(data, g[i]) // The points.
		}
		return data

	case geo.MultiLineString:
		data = append(data, GeoMultiLineString)
		data = append(data, len(g)) // Number of lines in the Multi line string
		for _, l := range g {
			ld := encode(bom, l)
			if ld == nil {
				return nil
			}
			data = append(data, ld...)
		}
		return data

	case geo.Polygon:
		data = append(data, GeoPolygon)
		data = append(data, uint32(len(g))) // Number of rings in the polygon
		for _, r := range g {
			data = append(data, uint32(len(r))) // Number of points in the ring
			for _, p := range r {
				data = append(data, p) // The points in the ring
			}
		}
		return data
	case geo.MultiPolygon:
		data = append(data, GeoMultiPolygon)
		data = append(data, uint32(len(g))) // Number of Polygons in the Multi.
		for _, p := range g {
			pd := encode(bom, p)
			if pd == nil {
				return nil
			}
			data = append(data, pd...)
		}
		return data
	case geo.Collection:
		data = append(data, GeoGeometryCollection)
		data = append(data, uint32(len(g))) // Number of Geometries
		for _, geom := range g {
			gd := encode(bom, geom)
			if gd == nil {
				return nil
			}
			data = append(data, gd...)
		}
		return data
	}
}

// Encode will encode the given Geometry as a binary representation with the given
// byte order, and write it to the provided io.Writer.
func Encode(w io.Writer, bom binary.ByteOrder, geom geo.Geometry) error {
	data := encode(bom, geom)
	if data == nil {
		return fmt.Errorf("Unabled to encode %v", geom)
	}
	return binary.Write(w, bom, data)
}

// DecodeBytes will decode the type into a Geometry
func DecodeBytes(b []byte) (geo.Geometry, error) {
	buff := bytes.NewReader(b)
	return Decode(buff)
}

// Decode is the main function that given a io.Reader will attempt to decode the
// Geometry from the byte stream.
func Decode(r io.Reader) (geo.Geometry, error) {
	byteOrder, typ, err := decodeByteOrderType(r)

	if err != nil {
		return nil, err
	}

	var geom Geometry
	switch typ {
	case GeoPoint:
		geom = new(Point)
	case GeoMultiPoint:
		geom = new(MultiPoint)
	case GeoLineString:
		geom = new(LineString)
	case GeoMultiLineString:
		geom = new(MultiLineString)
	case GeoPolygon:
		geom = new(Polygon)
	case GeoMultiPolygon:
		geom = new(MultiPolygon)
	case GeoGeometryCollection:
		geom = new(Collection)
	default:
		return nil, fmt.Errorf("Unknown Geometry! %v", typ)
	}
	if err := geom.Decode(byteOrder, r); err != nil {
		return nil, err
	}
	return geom.Geometry(), nil
}
