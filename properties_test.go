package geojson

func propertiesTestFeature() *Feature {
	rawJSON := `
	  { "type": "Feature",
	    "geometry": {"type": "Point", "coordinates": [102.0, 0.5]},
	    "properties": {"bool":true,"falsebool":false,"int": 1,"float64": 1.2,"string":"text"}
	  }`

	f, _ := UnmarshalFeature([]byte(rawJSON))
	return f
}
