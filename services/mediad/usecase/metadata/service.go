/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package metadata

import (
	"ac9/glad/services/mediad/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateMetadata(url string, total int, contentType entity.ContentType) (*entity.Metadata, error) {
	m, err := entity.NewMetadata(url, total, contentType)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(m); err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Service) GetMetadata(contentType entity.ContentType) (*entity.Metadata, error) {
	m, err := s.repo.Get(contentType)
	if err != nil {
		return nil, err
	}

	return m, nil
}