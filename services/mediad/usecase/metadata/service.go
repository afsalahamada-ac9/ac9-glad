/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package metadata

import entity "ac9/glad/services/mediad/entity"

// Service metadata usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateMetadata(version int64, url string, total int, metadataType string) error {
	switch metadataType {
	case "quote":
		q, err := entity.NewQuote(version, url, total)
		if err != nil {
			return err
		}
		s.repo.CreateQuote(q)
	case "media":
		m, err := entity.NewMedia(version, url, total)
		if err != nil {
			return err
		}
		s.repo.CreateMedia(m)
	}
	return nil
}
