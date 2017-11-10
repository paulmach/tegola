package server

import (
	"fmt"

	"github.com/paulmach/geo"
	"github.com/paulmach/tegola"
	"github.com/paulmach/tegola/mvt"
)

//	creates a debug layer with z/x/y encoded as a point
func debugLayer(tile tegola.Tile) *mvt.Layer {
	//	get tile bounding box
	ext := tile.BoundingBox()

	//	create a new layer and name it
	layer := mvt.Layer{
		Name: "debug",
	}

	//	tile outlines
	outline := mvt.Feature{
		Tags: map[string]interface{}{
			"type": "debug_outline",
		},
		Geometry: geo.LineString(ext.ToRing()),
	}

	//	new feature
	zxy := mvt.Feature{
		Tags: map[string]interface{}{
			"type": "debug_text",
			"zxy":  fmt.Sprintf("Z:%v, X:%v, Y:%v", tile.Z, tile.X, tile.Y),
		},
		Geometry: ext.Center(), // middle of the tile
	}

	layer.AddFeatures(outline, zxy)

	return &layer
}
