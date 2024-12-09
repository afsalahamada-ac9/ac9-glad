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

// Device data
type Device struct {
	ID        id.ID
	TenantID  id.ID
	AccountID id.ID

	PushToken  string
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
	pushToken string,
	revokeID *string,
	appVersion string,
	deviceInfo *string,
	platformInfo *string,
) (*Device, error) {
	t := &Device{
		ID:           id.IDInvalid,
		TenantID:     tenantID,
		AccountID:    accountID,
		PushToken:    pushToken,
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
		l.Log.Errorf("device=%#v, id is invalid", d)
		return glad.ErrInvalidEntity
	}

	if d.PushToken == "" || d.AccountID == id.IDInvalid || d.TenantID == id.IDInvalid || d.AppVersion == "" {
		l.Log.Errorf("device=%#v, invalid values for mandatory fields", d)
		return glad.ErrInvalidEntity
	}

	return nil
}
