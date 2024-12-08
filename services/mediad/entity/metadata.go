/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"ac9/glad/pkg/glad"
	"time"
)

type Quote struct {
	Version     int64
	LastUpdated time.Time
	Total       int
	URL         string
	CreatedAt   time.Time
}

type Media struct {
	Version     int64
	LastUpdated time.Time
	Total       int
	URL         string
	CreatedAt   time.Time
}

// NewQuote creates a new Quote instance
func NewQuote(version int64, url string, total int) (*Quote, error) {
	q := &Quote{
		Version: version,
		URL:     url,
		Total:   total,
	}
	if err := q.Validate(); err != nil {
		return nil, err
	}
	return q, nil
}

// NewMedia creates a new Media instance
func NewMedia(version int64, url string, total int) (*Media, error) {
	m := &Media{
		Version: version,
		URL:     url,
		Total:   total,
	}
	if err := m.Validate(); err != nil {
		return nil, err
	}
	return m, nil
}

// Validate validates the Quote fields
func (q *Quote) Validate() error {
	if q.URL == "" || q.Version < 0 {
		return glad.ErrInvalidEntity
	}
	return nil
}

// Validate validates the Media fields
func (m *Media) Validate() error {
	if m.URL == "" || m.Version < 0 {
		return glad.ErrInvalidEntity
	}
	return nil
}
