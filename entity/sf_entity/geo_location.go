package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type GeoLocation struct {
	Lat  float64 `json:"Geolocation__Latitude__s"`
	Long float64 `json:"Geolocation__Longitude__s"`
}

func (g GeoLocation) Value() (driver.Value, error) {
	return json.Marshal(g)
}

func (g *GeoLocation) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("invalid scan source for geolocation")
	}
	return json.Unmarshal(bytes, &g)
}
