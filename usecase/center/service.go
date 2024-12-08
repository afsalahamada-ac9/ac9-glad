/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package center

import (
	"strings"
	"time"

	"ac9/glad/entity"
	"ac9/glad/pkg/id"
)

// Service center usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateCenter creates a center
func (s *Service) CreateCenter(tenantID id.ID,
	extID,
	extName,
	name string,
	mode entity.CenterMode,
	isEnabled bool,
) (id.ID, error) {
	c, err := entity.NewCenter(tenantID, extID, extName, name, entity.CenterAddress{},
		entity.CenterGeoLocation{}, 0, mode, "", false, isEnabled)
	if err != nil {
		return id.IDInvalid, err
	}
	return s.repo.Create(c)
}

// GetCenter retrieves a center
func (s *Service) GetCenter(id id.ID) (*entity.Center, error) {
	t, err := s.repo.Get(id)
	if t == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return t, nil
}

// SearchCenters search center
func (s *Service) SearchCenters(tenantID id.ID,
	query string, page, limit int,
) ([]*entity.Center, error) {
	centers, err := s.repo.Search(tenantID, strings.ToLower(query), page, limit)
	if err != nil {
		return nil, err
	}
	if len(centers) == 0 {
		return nil, entity.ErrNotFound
	}
	return centers, nil
}

// ListCenters list center
func (s *Service) ListCenters(tenantID id.ID, page, limit int) ([]*entity.Center, error) {
	centers, err := s.repo.List(tenantID, page, limit)
	if err != nil {
		return nil, err
	}
	if len(centers) == 0 {
		return nil, entity.ErrNotFound
	}
	return centers, nil
}

// DeleteCenter Delete a center
func (s *Service) DeleteCenter(id id.ID) error {
	t, err := s.GetCenter(id)
	if t == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// UpdateCenter Update a center
func (s *Service) UpdateCenter(c *entity.Center) error {
	err := c.Validate()
	if err != nil {
		return err
	}
	c.UpdatedAt = time.Now()
	return s.repo.Update(c)
}

// GetCount gets total center count
func (s *Service) GetCount(tenantID id.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}
