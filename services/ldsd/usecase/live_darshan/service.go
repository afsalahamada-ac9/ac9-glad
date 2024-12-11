/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package livedarshan

import (
	"ac9/glad/pkg/id"
	"ac9/glad/services/ldsd/entity"
	"time"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return Service{
		repo: r,
	}
}

func (s *Service) CreateLiveDarshan(id int64, date string, startTime time.Time, meetingUrl string, createdBy id.ID) (*entity.LiveDarshan, error) {
	ld, err := entity.NewLiveDarshan(id, date, startTime, meetingUrl, createdBy)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ld); err != nil {
		return nil, err
	}

	return ld, nil
}

func (s *Service) GetLiveDarshan(id int64) (*entity.LiveDarshan, error) {
	ld, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return ld, nil
}

func (s *Service) GetAllLiveDarshan() ([]*entity.LiveDarshan, error) {
	ld, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return ld, nil
}
