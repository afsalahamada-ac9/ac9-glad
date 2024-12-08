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

// NewDevice create a new device without id i.e., invalid id
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
		ID:           id.IDInvalid,
		TenantID:     tenantID,
		AccountID:    accountID,
		Token:        token,
		RevokeID:     revokeID,
		AppVersion:   appVersion,
		DeviceInfo:   deviceInfo,
		PlatformInfo: platformInfo,
	}
	err := t.Validate(false)
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}
	return t, nil
}

// New creates a new device from existing device and overrides id and created & updated date
func (c Device) New() (*Device, error) {
	device := &c

	device.ID = id.New()
	device.CreatedAt = time.Now()
	device.UpdatedAt = device.CreatedAt

	err := device.Validate(true)
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}
	return device, nil
}

// Validate validates device
func (d *Device) Validate(isID bool) error {
	if isID && d.ID == id.IDInvalid {
		return glad.ErrInvalidEntity
	}

	if d.Token == "" || d.AccountID == id.IDInvalid || d.TenantID == id.IDInvalid || d.AppVersion == "" {
		return glad.ErrInvalidEntity
	}

	return nil
}
