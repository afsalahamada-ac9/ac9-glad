/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package device

import (
	"ac9/glad/pkg/id"
	"ac9/glad/services/pushd/entity"
	"context"
)

// DeviceReader device reader
type DeviceReader interface {
	GetByAccount(tenantID id.ID, accountID id.ID) ([]*entity.Device, error)
	GetCount(tenantID id.ID) (int, error)
}

// DeviceWriter device writer
type DeviceWriter interface {
	Create(e *entity.Device) (id.ID, error)
	Delete(id id.ID) error
}

// Device repository interface
type DeviceRepository interface {
	DeviceReader
	DeviceWriter
}

// Push notification service interface
type PushNS interface {
	Send(ctx context.Context, token string, header, content string) error
}

// UseCase interface
type UseCase interface {
	Create(device entity.Device) (id.ID, error)
	GetByAccount(tenantID id.ID, accountID id.ID) ([]*entity.Device, error)
	Delete(id id.ID) error
	GetCount(id id.ID) int

	// returns http status codes. if notification is sent successfully to one device
	// in the account, then push notification for that account is considered as successful
	// 201 - successfully pushed the notification
	// 404 - token not found
	// 500 - unknown error
	Notify(tenantID id.ID, accountID []id.ID, header, content string) ([]int, error)
}
