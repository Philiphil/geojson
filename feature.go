package geojson

import (
	"encoding/json"
)

// A Feature corresponds to GeoJSON feature object
type Feature struct {
	ID          interface{}            `json:"id,omitempty"`
	Type        string                 `json:"type"`
	BoundingBox []float64              `json:"bbox,omitempty"`
	Geometry    *Geometry              `json:"geometry"`
	Properties  Properties `json:"properties"`
	CRS         map[string]interface{} `json:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
}

// NewFeature creates and initializes a GeoJSON feature given the required attributes.
func NewFeature(geometry *Geometry) *Feature {
	return &Feature{
		Type:       "Feature",
		Geometry:   geometry,
	}
}

// NewPointFeature creates and initializes a GeoJSON feature with a point geometry using the given coordinate.
func NewPointFeature(coordinate []float64) *Feature {
	return NewFeature(NewPointGeometry(coordinate))
}

// NewMultiPointFeature creates and initializes a GeoJSON feature with a multi-point geometry using the given coordinates.
func NewMultiPointFeature(coordinates ...[]float64) *Feature {
	return NewFeature(NewMultiPointGeometry(coordinates...))
}

// NewLineStringFeature creates and initializes a GeoJSON feature with a line string geometry using the given coordinates.
func NewLineStringFeature(coordinates [][]float64) *Feature {
	return NewFeature(NewLineStringGeometry(coordinates))
}

// NewMultiLineStringFeature creates and initializes a GeoJSON feature with a multi-line string geometry using the given lines.
func NewMultiLineStringFeature(lines ...[][]float64) *Feature {
	return NewFeature(NewMultiLineStringGeometry(lines...))
}

// NewPolygonFeature creates and initializes a GeoJSON feature with a polygon geometry using the given polygon.
func NewPolygonFeature(polygon [][][]float64) *Feature {
	return NewFeature(NewPolygonGeometry(polygon))
}

// NewMultiPolygonFeature creates and initializes a GeoJSON feature with a multi-polygon geometry using the given polygons.
func NewMultiPolygonFeature(polygons ...[][][]float64) *Feature {
	return NewFeature(NewMultiPolygonGeometry(polygons...))
}

// NewCollectionFeature creates and initializes a GeoJSON feature with a geometry collection geometry using the given geometries.
func NewCollectionFeature(geometries ...*Geometry) *Feature {
	return NewFeature(NewCollectionGeometry(geometries...))
}

// MarshalJSON converts the feature object into the proper JSON.
// It will handle the encoding of all the child geometries.
// Alternately one can call json.Marshal(f) directly for the same result.
func (f Feature) MarshalJSON() ([]byte, error) {
	type feature Feature

	fea := &feature{
		ID:       f.ID,
		Type:     "Feature",
		Geometry: f.Geometry,
		Properties: f.Properties,
	}

	if f.BoundingBox != nil && len(f.BoundingBox) != 0 {
		fea.BoundingBox = f.BoundingBox
	}

	if f.CRS != nil && len(f.CRS) != 0 {
		fea.CRS = f.CRS
	}

	return json.Marshal(fea)
}

// UnmarshalFeature decodes the data into a GeoJSON feature.
// Alternately one can call json.Unmarshal(f) directly for the same result.
func UnmarshalFeature(data []byte) (*Feature, error) {
	f := &Feature{}
	err := json.Unmarshal(data, f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
