/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"fmt"
	"time"
)

type ContentType string

const (
	MediaQuote ContentType = "quote"
	MediaImage ContentType = "image"
)

type Metadata struct {
	Version      int64
	URL          string
	Total        int
	LastUpdated  time.Time
	CreatedAt    time.Time
	Type ContentType
}

func NewMetadata(url string, total int, contentType ContentType) (*Metadata, error) {
	m := &Metadata{
		URL:          url,
		Total:        total,
		Type: contentType,
	}

	if err := m.validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return m, nil
}

func (m *Metadata) validate() error {
	if m.URL == "" {
		return fmt.Errorf("URL cannot be empty")
	}
	return nil
}

