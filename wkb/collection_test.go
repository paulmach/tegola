package wkb_test

import (
	"encoding/binary"
	"testing"

	"github.com/paulmach/geo"
	"github.com/paulmach/tegola/wkb"
)

func newPoint(x, y float64) *wkb.Point {
	p := wkb.NewPoint(x, y)
	return &p
}
func TestCollection(t *testing.T) {
	testcases := TestCases{
		{
			bytes: []byte{
				//01    02    03    04    05    06    07    08
				0x02, 0x00, 0x00, 0x00, // Number of Geometries in Collection
				0x01,                   // Byte order marker little
				0x01, 0x00, 0x00, 0x00, // Type (1) Point
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x40, // X1 4
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x18, 0x40, // Y1 6
				0x01,                   // Byte order marker little
				0x02, 0x00, 0x00, 0x00, // Type (2) Line
				0x02, 0x00, 0x00, 0x00, // Number of Points (2)
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x40, // X1 4
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x18, 0x40, // Y1 6
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1c, 0x40, // X2 7
				0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x24, 0x40, // Y2 10
			},
			bom: binary.LittleEndian,
			expected: &wkb.Collection{
				Collection: geo.Collection{
					geo.Point{4, 6},
					geo.LineString{{4, 6}, {7, 10}},
				},
			},
		},
	}
	testcases.RunTests(t, func(num int, tcase *TestCase) {
		var p, expected wkb.Collection
		if cexp, ok := tcase.expected.(*wkb.Collection); !ok {
			t.Errorf("Bad test case %v", num)
			return
		} else {
			expected = *cexp
		}
		if err := p.Decode(tcase.bom, tcase.Reader()); err != nil {
			t.Errorf("Got unexpected error %v", err)
			return
		}
		if len(expected.Collection) != len(p.Collection) {
			t.Errorf("Test %v: Collection did not get decoded correctly, expected: %v got: %v", num, expected, p)
			return
		}
		for i := range expected.Collection {
			if expected.Collection[i].GeoJSONType() != p.Collection[i].GeoJSONType() {
				t.Errorf("Test %v: expected[%v]: %v got: %v", num, i, expected, p)
			}
		}

	})

}
