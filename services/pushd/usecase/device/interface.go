/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package device

import (
	"ac9/glad/pkg/id"
	"ac9/glad/services/pushd/entity"
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

// UseCase interface
type UseCase interface {
	Create(device entity.Device) (id.ID, error)
	GetByAccount(tenantID id.ID, accountID id.ID) ([]*entity.Device, error)
	Delete(id id.ID) error
	GetCount(id id.ID) int
}
