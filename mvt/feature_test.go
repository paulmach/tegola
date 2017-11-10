package mvt

import (
	"testing"

	"github.com/paulmach/geo"
	"github.com/paulmach/tegola/mvt/vector_tile"
)

func TestEncodeGeometry(t *testing.T) {
	/*
		complexGemo := basic.Polygon{
			basic.Line{
				basic.Point{8, 8.5},
				basic.Point{9, 9},
				basic.Point{20, 20},
				basic.Point{11, 20},
			},
		}
	*/
	testcases := []struct {
		geo  geo.Geometry
		typ  vectorTile.Tile_GeomType
		bbox geo.Bound
		egeo []uint32
		eerr error
	}{
		{
			geo:  nil,
			typ:  vectorTile.Tile_UNKNOWN,
			bbox: geo.NewBound(0, 4096, 0, 4096),
			egeo: []uint32{},
			eerr: ErrNilGeometryType,
		},
		{
			geo:  geo.Point{1, 1},
			typ:  vectorTile.Tile_POINT,
			bbox: geo.NewBound(0, 4096, 0, 4096),
			egeo: []uint32{9, 2, 2},
		},
		{
			geo:  geo.Point{25, 17},
			typ:  vectorTile.Tile_POINT,
			bbox: geo.NewBound(0, 4096, 0, 4096),
			egeo: []uint32{9, 50, 34},
		},
		{
			geo:  geo.MultiPoint{{5, 7}, {3, 2}},
			typ:  vectorTile.Tile_POINT,
			bbox: geo.NewBound(0, 4096, 0, 4096),
			egeo: []uint32{17, 10, 14, 3, 9},
		},
		{
			geo:  geo.LineString{{2, 2}, {2, 10}, {10, 10}},
			typ:  vectorTile.Tile_LINESTRING,
			bbox: geo.NewBound(0, 4096, 0, 4096),
			egeo: []uint32{9, 4, 4, 18, 0, 16, 16, 0},
		},
		{
			geo: geo.MultiLineString{
				{{2, 2}, {2, 10}, {10, 10}},
				{{1, 1}, {3, 5}},
			},
			typ:  vectorTile.Tile_LINESTRING,
			bbox: geo.NewBound(0, 4096, 0, 4096),
			egeo: []uint32{9, 4, 4, 18, 0, 16, 16, 0, 9, 17, 17, 10, 4, 8},
		},
		{
			geo: geo.Polygon{
				{{3, 6}, {8, 12}, {20, 34}},
			},
			typ:  vectorTile.Tile_POLYGON,
			bbox: geo.NewBound(0, 4096, 0, 4096),
			egeo: []uint32{9, 6, 12, 18, 10, 12, 24, 44, 15},
		},
		{
			geo: geo.MultiPolygon{
				{
					{{0, 0}, {10, 0}, {10, 10}, {0, 10}},
				},
				{
					{{11, 11}, {20, 11}, {20, 20}, {11, 20}},
					{{13, 13}, {13, 17}, {17, 17}, {17, 13}},
				},
			},
			typ:  vectorTile.Tile_POLYGON,
			bbox: geo.NewBound(0, 4096, 0, 4096),
			egeo: []uint32{9, 0, 0, 26, 20, 0, 0, 20, 19, 0, 15, 9, 22, 2, 26, 18, 0, 0, 18, 17, 0, 15, 9, 4, 13, 26, 0, 8, 8, 0, 0, 7, 15},
		},
	}
	for _, tcase := range testcases {
		g, gtype, err := encodeGeometry(tcase.geo, tcase.bbox, 4096)
		if tcase.eerr != err {
			t.Errorf("Expected error (%v) got (%v) instead", tcase.eerr, err)
		}
		if gtype != tcase.typ {
			t.Errorf("Expected Geometry Type to be %v Got: %v", tcase.typ, gtype)
		}
		if len(g) != len(tcase.egeo) {
			t.Errorf("Geometry length is not what was expected(%v) got (%v)", tcase.egeo, g)
			continue
		}
		for i := range tcase.egeo {
			if tcase.egeo[i] != g[i] {
				t.Errorf("Geometry is not what was expected(%v) got (%v)", tcase.egeo, g)
				break
			}
		}
	}
}

func TestNewFeature(t *testing.T) {
	testcases := []struct {
		geo      geo.Geometry
		tags     map[string]interface{}
		expected []Feature
	}{
		{
			geo:      nil,
			tags:     nil,
			expected: []Feature{},
		},
	}
	for i, tcase := range testcases {
		got := NewFeatures(tcase.geo, tcase.tags)
		if len(tcase.expected) != len(got) {
			t.Errorf("Test %v: Expected to get %v features got %v features.", i, len(tcase.expected), len(got))
			continue
		}
		if len(tcase.expected) <= 0 {
			continue
		}
		// TODO test to make sure we got the correct feature

	}
}

func TestNormalizePoint(t *testing.T) {
	testcases := []struct {
		point       geo.Point
		bbox        geo.Bound
		nx, ny      int64
		layerExtent int
	}{
		{
			point:       geo.Point{960000, 6002729},
			bbox:        geo.NewBound(958826.08, 978393.96, 5987771.04, 6007338.92),
			nx:          245,
			ny:          3131,
			layerExtent: 4096,
		},
	}

	for i, tcase := range testcases {
		//	new cursor
		c := newCursor(tcase.bbox, tcase.layerExtent)

		nx, ny := c.ScalePoint(tcase.point)
		if nx != tcase.nx {
			t.Errorf("Test %v: Expected nx value of %v got %v.", i, tcase.nx, nx)
		}
		if ny != tcase.ny {
			t.Errorf("Test %v: Expected ny value of %v got %v.", i, tcase.ny, ny)
		}
		continue
	}
}
