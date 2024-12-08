/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package device

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	"ac9/glad/services/pushd/entity"
)

// inmemDevice in memory repo
type inmemDevice struct {
	m map[id.ID]*entity.Device
}

// newinmemDevice create new repository
func newInmemDevice() *inmemDevice {
	var m = map[id.ID]*entity.Device{}
	return &inmemDevice{
		m: m,
	}
}

// Create a device
func (r *inmemDevice) Create(e *entity.Device) (id.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// GetByAccount retrieve device(s) by account id
func (r *inmemDevice) GetByAccount(tenantID id.ID,
	accountID id.ID,
) ([]*entity.Device, error) {
	var devices []*entity.Device
	for _, j := range r.m {
		if j.TenantID == tenantID &&
			j.AccountID == accountID {
			devices = append(devices, j)
		}
	}

	return devices, nil
}

// Delete a device
func (r *inmemDevice) Delete(id id.ID) error {
	if r.m[id] == nil {
		return glad.ErrNotFound
	}
	r.m[id] = nil
	delete(r.m, id)
	return nil
}

// GetCount gets total devices for a given tenant
func (r *inmemDevice) GetCount(tenantID id.ID) (int, error) {
	count := 0
	for _, j := range r.m {
		if j.TenantID == tenantID {
			count++
		}
	}
	return count, nil
}
