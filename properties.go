package geojson

type Address struct {
	Tourism     string
	Road        string
	Suburb      string
	City        string
	County      string
	State       string
	Postcode    string
	Country     string
	CountryCode string `json:"country_code"`
}

type Properties struct {
	PlaceID     int    `json:"place_id"`
	OsmType     string `json:"osm_type"`
	OsmID       int    `json:"osm_id"`
	DisplayName string `json:"display_name"`
	PlaceRank   int    `json:"place_rank"`
	Category    string
	Type        string
	Importance  float64
	Icon        *string
	AddressType *string `json:"addresstype"`
	Name        *string
	Address     *Address
}
