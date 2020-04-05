package t38c

import geojson "github.com/paulmach/go.geojson"

// SearchArea ...
type SearchArea Command

// AreaGet ...
func AreaGet(objectID string) SearchArea {
	return SearchArea(NewCommand("GET", objectID))
}

// AreaBounds ...
func AreaBounds(minlat, minlon, maxlat, maxlon float64) SearchArea {
	return SearchArea(NewCommand("BOUNDS", minlat, minlon, maxlat, maxlon))
}

// AreaFeatureCollection ...
func AreaFeatureCollection(fc *geojson.FeatureCollection) SearchArea {
	// TODO: handle error?
	b, _ := fc.MarshalJSON()
	return SearchArea(NewCommand("OBJECT", string(b)))
}

// AreaFeature ...
func AreaFeature(ft *geojson.Feature) SearchArea {
	// TODO: handle error?
	b, _ := ft.MarshalJSON()
	return SearchArea(NewCommand("OBJECT", string(b)))
}

// AreaGeometry ...
func AreaGeometry(gm *geojson.Geometry) SearchArea {
	// TODO: handle error?
	b, _ := gm.MarshalJSON()
	return SearchArea(NewCommand("OBJECT", string(b)))

}

// AreaCircle ...
func AreaCircle(lat, lon, meters float64) SearchArea {
	return SearchArea(NewCommand("CIRCLE", lat, lon, meters))
}

// NearbyArea ...
type NearbyArea Command

// NearbyPoint ...
func NearbyPoint(lat, lon, meters float64) NearbyArea {
	return NearbyArea(NewCommand("POINT", lat, lon, meters))
}

// SetArea ...
type SetArea Command

// SetPoint ...
func SetPoint(lat, lon float64) SetArea {
	return SetArea(NewCommand("POINT", lat, lon))
}

// SetPointZ ...
func SetPointZ(lat, lon, z float64) SetArea {
	return SetArea(NewCommand("POINT", lat, lon, z))
}

// SetFeatureCollection ...
func SetFeatureCollection(fc *geojson.FeatureCollection) SetArea {
	b, _ := fc.MarshalJSON()
	return SetArea(NewCommand("OBJECT", string(b)))
}

// SetFeature ...
func SetFeature(ft *geojson.Feature) SetArea {
	b, _ := ft.MarshalJSON()
	return SetArea(NewCommand("OBJECT", string(b)))
}

// SetGeometry ...
func SetGeometry(gm *geojson.Geometry) SetArea {
	b, _ := gm.MarshalJSON()
	return SetArea(NewCommand("OBJECT", string(b)))
}