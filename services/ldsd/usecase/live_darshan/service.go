/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package live_darshan

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	"ac9/glad/services/ldsd/entity"
	"database/sql"
)

// Service live darshan usecase
type Service struct {
	repo Repository
}

// NewService creates new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateLiveDarshan creates a live darshan
func (s *Service) CreateLiveDarshan(
	tenantID id.ID,
	date string,
	startTime string,
	meetingURL string,
	createdBy id.ID,
) (*entity.LiveDarshan, error) {
	ld, err := entity.NewLiveDarshan(tenantID, date, startTime, meetingURL, createdBy)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ld); err != nil {
		return nil, err
	}

	return ld, nil
}

// GetLiveDarshan retrieves a live darshan
func (s *Service) GetLiveDarshan(id int64) (*entity.LiveDarshan, error) {
	ld, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return ld, nil
}

// ListLiveDarshan lists live darshan
func (s *Service) ListLiveDarshan(tenantID id.ID, page, limit int) ([]*entity.LiveDarshan, error) {
	ld, err := s.repo.List(tenantID, page, limit)
	if err != nil {
		return nil, err
	}

	return ld, nil
}

// DeleteLiveDarshan deletes a live darshan
func (s *Service) DeleteLiveDarshan(ldID int64) error {
	err := s.repo.Delete(ldID)
	if err == sql.ErrNoRows {
		return glad.ErrNotFound
	}

	return err
}

// GetCount gets total live darshan count
func (s *Service) GetCount(tenantID id.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}
