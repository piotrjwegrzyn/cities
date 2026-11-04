package crawler

type Schema struct {
	Geonames []Geoname `json:"geonames"`
}

type Geoname struct {
	Name string `json:"name"`
}
