package wkb

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/paulmach/geo"
)

// LineString describes a line, that is made up of two or more points
type LineString struct {
	geo.LineString
}

// Type returns the type constant for a LineString
func (LineString) Type() uint32 {
	return GeoLineString
}

func (ls *LineString) Geometry() geo.Geometry {
	return ls.LineString
}

// Decode will decode the binary representation into a LineString Object.
func (ls *LineString) Decode(bom binary.ByteOrder, r io.Reader) error {
	var num uint32
	if err := binary.Read(r, bom, &num); err != nil {
		return err
	}
	for i := 0; i < int(num); i++ {
		var p = new(Point)
		if err := p.Decode(bom, r); err != nil {
			return err
		}
		ls.LineString = append(ls.LineString, p.Point)
	}
	return nil
}

//String returns the WKT representation of the Geometry
func (ls *LineString) String() string {
	return WKT(ls) // If we have a failure we don't care
}

//MultiLineString represents one or more independent lines.
type MultiLineString struct {
	geo.MultiLineString
}

//Type returns the Type constant for a Multiline String.
func (MultiLineString) Type() uint32 {
	return GeoMultiLineString
}

func (ml *MultiLineString) Geometry() geo.Geometry {
	return ml.MultiLineString
}

//Decode takes a byteOrder and an io.Reader, to decode the stream.
func (ml *MultiLineString) Decode(bom binary.ByteOrder, r io.Reader) error {
	var num uint32
	if err := binary.Read(r, bom, &num); err != nil {
		return err
	}
	for i := uint32(0); i < num; i++ {
		var l = new(LineString)
		byteOrder, typ, err := decodeByteOrderType(r)
		if err != nil {
			return err
		}
		if typ != GeoLineString {
			return fmt.Errorf("Expect Multilines to contains lines; did not find a line.")
		}
		if err := l.Decode(byteOrder, r); err != nil {
			return err
		}
		ml.MultiLineString = append(ml.MultiLineString, l.LineString)
	}
	return nil
}

//String returns the WKT representation of the Geometry.
func (ml *MultiLineString) String() string {
	return WKT(ml.MultiLineString)
}
