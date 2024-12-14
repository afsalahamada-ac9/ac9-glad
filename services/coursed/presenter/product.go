/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/id"
	"time"

	l "ac9/glad/pkg/logger"

	"github.com/ulule/deepcopier"
)

// Product data - TenantID is returned in the HTTP header
// X-GLAD-TenantID
type ProductReq struct {
	ExtName          string                   `json:"extName"`
	Title            string                   `json:"title"`
	CType            string                   `json:"ctype"`
	BaseProductExtID string                   `json:"baseProductExtID"`
	DurationDays     int32                    `json:"durationDays"`
	Visibility       entity.ProductVisibility `json:"visibility"`
	MaxAttendees     int32                    `json:"maxAttendees"`
	Format           entity.ProductFormat     `json:"format"`
	IsAutoApprove    bool                     `json:"isAutoApprove"`
}

type ProductResponse struct {
	ID id.ID `json:"id"`
}

type Product struct {
	ProductReq
	ProductResponse
}

type ProductFull struct {
	ExtID string `json:"extID"`
	ProductReq
	ProductResponse
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// FromEntityProduct creates presenter product from entity
func (p *Product) FromEntityProduct(e *entity.Product) error {
	p.ID = e.ID
	p.ExtName = e.ExtName
	p.Title = e.Title
	p.CType = e.CType
	p.BaseProductExtID = e.BaseProductExtID
	p.DurationDays = e.DurationDays
	p.Visibility = e.Visibility
	p.MaxAttendees = e.MaxAttendees
	p.Format = e.Format
	p.IsAutoApprove = e.IsAutoApprove

	return nil
}

func (pf ProductFull) ToEntity(e *entity.Product) error {
	deepcopier.Copy(pf).To(e)
	l.Log.Infof("Product full=%v, product=%v", pf, e)
	return nil
}
