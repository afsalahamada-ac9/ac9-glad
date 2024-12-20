/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"time"
)

// Center mode
type CenterMode string

const (
	CenterInPerson CenterMode = "in-person"
	CenterOnline   CenterMode = "online"
	CenterNotSet   CenterMode = "not-set"
	// Add new types here
)

// Center Address
type CenterAddress struct {
	Street1 string
	Street2 string
	City    string
	State   string
	Zip     string
	Country string
}

// Center GeoLocation
type CenterGeoLocation struct {
	Lat  float64
	Long float64
}

// Center data
type Center struct {
	ID       id.ID
	TenantID id.ID
	ExtID    string

	ExtName     string
	Name        string
	Address     CenterAddress
	GeoLocation CenterGeoLocation

	Capacity int32
	Mode     CenterMode

	WebPage    string
	IsNational bool
	IsEnabled  bool

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewCenterAddress creates a new center address
func NewCenterAddress(street1 string,
	street2 string,
	city string,
	state string,
	zip string,
	country string) (*CenterAddress, error) {

	l := &CenterAddress{
		Street1: street1,
		Street2: street2,
		City:    city,
		State:   state,
		Zip:     zip,
		Country: country,
	}
	err := l.Validate()
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}
	return l, nil
}

// Validate validates center address
func (l *CenterAddress) Validate() error {
	if l.Street1 == "" || l.City == "" || l.State == "" || l.Zip == "" || l.Country == "" {
		return glad.ErrInvalidEntity
	}
	return nil
}

// NewCenterGeoLocation creates a new center geo location
func NewCenterGeoLocation(lat float64, long float64) (*CenterGeoLocation, error) {
	g := &CenterGeoLocation{
		Lat:  lat,
		Long: long,
	}
	err := g.Validate()
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}
	return g, nil
}

// Validate validates center geo location
func (g *CenterGeoLocation) Validate() error {
	if g.Lat == 0 || g.Long == 0 {
		return glad.ErrInvalidEntity
	}
	return nil
}

// NewCenter create a new center
func NewCenter(tenantID id.ID,
	name string,
	address CenterAddress,
	geoLocation CenterGeoLocation,
	capacity int32,
	mode CenterMode,
	webPage string,
	isNational bool,
	isEnabled bool,
) (*Center, error) {
	c := &Center{
		ID:          id.New(),
		TenantID:    tenantID,
		Name:        name,
		Address:     address,
		GeoLocation: geoLocation,
		Capacity:    capacity,
		Mode:        mode,
		WebPage:     webPage,
		IsNational:  isNational,
		IsEnabled:   isEnabled,
		CreatedAt:   time.Now(),
	}
	err := c.Validate()
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}
	return c, nil
}

// Transform fixes the data issues, if any
func (c *Center) Transform() {
	if c.Name == "" {
		// TODO: count the centers with empty name (center name in SF)
		c.Name = c.ExtName
	}

	if c.Mode == "" {
		// TODO: count the centers with empty mode (center mode in SF)
		c.Mode = CenterNotSet
	}
}

func (c *Center) Validate() error {
	if c.TenantID == id.IDInvalid {
		l.Log.Warnf("Invalid tenant id=%v, center extID=%v", c.TenantID, c.ExtID)
		return glad.ErrInvalidEntity
	}

	// When external id is present then name must be present
	if c.ExtID != "" && c.ExtName == "" {
		l.Log.Warnf("Center extID=%v has empty extName", c.ExtID)
		return glad.ErrInvalidEntity
	}

	if c.Name == "" {
		l.Log.Warnf("Center extID=%v has empty name", c.ExtID)
		return glad.ErrInvalidEntity
	}

	return nil
}
