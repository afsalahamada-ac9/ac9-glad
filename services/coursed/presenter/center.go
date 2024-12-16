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
	return nil
}

// GladCenterToEntity populates entity center from glad entity
func GladCenterToEntity(gp glad.Center, e *entity.Center) error {
	deepcopier.Copy(gp).To(e)
	deepcopier.Copy(gp.Address).To(&e.Address)
	e.Mode = entity.CenterMode(gp.Mode)

	// deepcopier is unable to copy these, as the names are different
	// Note: Can add field tags and try it
	e.GeoLocation.Lat = gp.GeoLocation.Latitude
	e.GeoLocation.Long = gp.GeoLocation.Longitude

	l.Log.Debugf("Center=%v, entity.center=%v", gp, e)
	return nil
}
