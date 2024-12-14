/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	"time"
)

// ProductVisibility represents the visibility status of a product
type ProductVisibility string

const (
	ProductVisibilityPublic   ProductVisibility = "Public"
	ProductVisibilityUnlisted ProductVisibility = "Unlisted"
)

// ProductFormat represents the format of a product
type ProductFormat string

const (
	ProductFormatInPerson    ProductFormat = "In Person"
	ProductFormatOnline      ProductFormat = "Online"
	ProductFormatDestination ProductFormat = "Destination Retreats"
)

// Product represents a product entity
type Product struct {
	ID               id.ID
	ExtID            string
	TenantID         id.ID
	ExtName          string
	Title            string
	CType            string
	BaseProductExtID string
	DurationDays     int32
	Visibility       ProductVisibility
	MaxAttendees     int32
	Format           ProductFormat

	IsAutoApprove bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewProduct creates a new product with the given parameters
func NewProduct(tenantID id.ID,
	extName string,
	title string,
	ctype string,
	baseProductExtID string,
	durationDays int32,
	visibility ProductVisibility,
	maxAttendees int32,
	format ProductFormat,
	isAutoApprove bool,
) (*Product, error) {
	p := &Product{
		ID:               id.New(),
		TenantID:         tenantID,
		ExtName:          extName,
		Title:            title,
		CType:            ctype,
		BaseProductExtID: baseProductExtID,
		DurationDays:     durationDays,
		Visibility:       visibility,
		MaxAttendees:     maxAttendees,
		Format:           format,
		IsAutoApprove:    isAutoApprove,
		CreatedAt:        time.Now(),
	}

	err := p.Validate()
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}

	return p, nil
}

// Validate validates the product fields
func (p *Product) Validate() error {
	if p.TenantID == 0 {
		return glad.ErrInvalidEntity
	}

	if p.ExtName == "" || p.Title == "" || p.CType == "" || p.BaseProductExtID == "" {
		return glad.ErrInvalidEntity
	}

	return nil
}
