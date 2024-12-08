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

// Device data
type Device struct {
	ID        id.ID
	TenantID  id.ID
	AccountID id.ID

	Token      string
	RevokeID   *string
	AppVersion string

	DeviceInfo   *string
	PlatformInfo *string

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewDevice create a new device
func NewDevice(
	tenantID id.ID,
	accountID id.ID,
	token string,
	revokeID *string,
	appVersion string,
	deviceInfo *string,
	platformInfo *string,
) (*Device, error) {
	t := &Device{
		ID:           id.New(),
		TenantID:     tenantID,
		AccountID:    accountID,
		Token:        token,
		RevokeID:     revokeID,
		AppVersion:   appVersion,
		DeviceInfo:   deviceInfo,
		PlatformInfo: platformInfo,
	}
	err := t.Validate()
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}
	return t, nil
}

// Validate validate device
func (t *Device) Validate() error {
	if t.Token == "" || t.AccountID == id.IDInvalid || t.TenantID == id.IDInvalid || t.AppVersion == "" {
		return glad.ErrInvalidEntity
	}

	return nil
}
