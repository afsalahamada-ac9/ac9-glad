/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"

	"github.com/ulule/deepcopier"
)

type CenterAddress struct {
	Street1    string `json:"street1,omitempty"`
	Street2    string `json:"street2,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"zip,omitempty"`
	Country    string `json:"country,omitempty"`
}

type CenterGeoLocation struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}
type CenterReq struct {
	Name string `json:"name"`
	// deprecated
	ExtName string `json:"extName"`
	// Address     CenterAddress     `json:"address,omitempty"`
	// GeoLocation CenterGeoLocation `json:"geoLocation,omitempty"`
	Capacity   int32             `json:"capacity"`
	Mode       entity.CenterMode `json:"mode"`
	IsNational bool              `json:"isNational"`
	IsEnabled  bool              `json:"isEnabled"`
	WebPage    string            `json:"webPage,omitempty"`
}

type CenterResponse struct {
	ID id.ID `json:"id"`
}

type Center struct {
	CenterReq
	CenterResponse
}

type CenterImportResponse struct {
	ID      id.ID  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}

// FromEntityCenter populates center struct from center entity
func (c *Center) FromEntityCenter(e *entity.Center) error {

	deepcopier.Copy(e).To(c)

	// c.ID = e.ID
	// c.ExtName = e.ExtName
	// c.Name = e.Name
	// c.Address1 = e.Address.Street1
	// c.Address2 = e.Address.Street2
	// c.City = e.Address.City
	// c.State = e.Address.State
	// c.PostalCode = e.Address.Zip
	// c.Country = e.Address.Country
	// c.Latitude = e.GeoLocation.Lat
	// c.Longitude = e.GeoLocation.Long
	// c.MaxCapacity = e.MaxCapacity
	// c.Mode = e.Mode
	// c.IsNational = e.IsNationalCenter
	// c.IsEnabled = e.IsEnabled
	// c.CenterURL = e.CenterURL
	return nil
}

func GladCenterToEntity(gp glad.Center, e *entity.Center) error {
	deepcopier.Copy(gp).To(e)
	deepcopier.Copy(gp.Address).To(&e.Address)
	e.Mode = entity.CenterMode(gp.Mode)
	e.GeoLocation.Lat = gp.GeoLocation.Latitude
	e.GeoLocation.Long = gp.GeoLocation.Longitude

	l.Log.Infof("Center=%v, entity.center=%v", gp, e)
	return nil
}
