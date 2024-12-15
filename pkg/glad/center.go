package glad

import (
	"time"
)

// Note: Ideally these should be proto files and we should use grpc between services
type CenterAddress struct {
	Street1    string `json:"street1"`
	Street2    string `json:"street2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"zip"`
	Country    string `json:"country"`
}

type CenterGeoLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}
type Center struct {
	ExtID       string `json:"extID"`
	ExtName     string `json:"extName"`
	Name        string `json:"name"`
	MaxCapacity int32  `json:"capacity"`
	Mode        string `json:"mode"`
	IsNational  bool   `json:"isNational"`
	IsEnabled   bool   `json:"isEnabled"`
	CenterURL   string `json:"webPage"`

	Address     CenterAddress     `json:"address"`
	GeoLocation CenterGeoLocation `json:"geoLocation"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CenterResponse struct {
	ID      int64  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}
