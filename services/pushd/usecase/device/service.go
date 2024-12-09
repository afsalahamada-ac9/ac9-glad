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

// Service device usecase
type Service struct {
	repo DeviceRepository
}

// NewService creates new service
func NewService(r DeviceRepository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateDevice creates a device
func (s *Service) Create(
	device entity.Device,
) (id.ID, error) {
	d, err := device.New()
	if err != nil {
		return id.IDInvalid, err
	}

	deviceID, err := s.repo.Create(d)
	if err != nil {
		return deviceID, err
	}
	return deviceID, err
}

// GetByAccount retrieves devices
func (s *Service) GetByAccount(tenantID id.ID, accountID id.ID) ([]*entity.Device, error) {
	t, err := s.repo.GetByAccount(tenantID, accountID)
	if t == nil {
		return nil, glad.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return t, nil
}

// DeleteDevice deletes a device
// Note: Since delete is cascaded to dependent tables, no need to call those functions explicitly
func (s *Service) Delete(id id.ID) error {
	return s.repo.Delete(id)
}

// GetCount gets total device count
func (s *Service) GetCount(tenantID id.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}
