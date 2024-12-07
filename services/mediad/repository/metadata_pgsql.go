/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"database/sql"

	mediadEntity "ac9/glad/services/mediad/entity"
)

// MetadataPGSQL postgres repo
type MetadataPGSQL struct {
	db *sql.DB
}

// NewMetadataPGSQL create new repository
func NewMetadataPGSQL(db *sql.DB) *MetadataPGSQL {
	return &MetadataPGSQL{
		db: db,
	}
}

// Create a quote
func (r *MetadataPGSQL) CreateQuote(e *mediadEntity.Quote) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO metadata (version, url, total, type) 
		VALUES($1, $2, $3, $4)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.Version,
		e.URL,
		e.Total,
		"quote",
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

// Create a media
func (r *MetadataPGSQL) CreateMedia(e *mediadEntity.Media) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO metadata (version, url, total, type) 
		VALUES($1, $2, $3, $4)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.Version,
		e.URL,
		e.Total,
		"media",
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
