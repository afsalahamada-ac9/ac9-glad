/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"ac9/glad/services/mediad/entity"
	"database/sql"
)

type MetadataPGSQL struct {
	db *sql.DB
}

func NewMetadataPGSQL(db *sql.DB) *MetadataPGSQL {
	return &MetadataPGSQL{
		db: db,
	}
}

func (r *MetadataPGSQL) Create(m *entity.Metadata) error {
	query := `INSERT INTO metadata (url, total, type) VALUES ($1, $2, $3)`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		m.URL,
		m.Total,
		m.Type,
	)
	return err
}

func (r *MetadataPGSQL) Get(contentType entity.ContentType) (*entity.Metadata, error) {
	query := `SELECT version, url, total, type, created_at, last_updated 
			  FROM metadata 
			  WHERE type = $1 
			  ORDER BY version DESC LIMIT 1`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(contentType)

	var m entity.Metadata
	err = row.Scan(
		&m.Version,
		&m.URL,
		&m.Total,
		&m.Type,
		&m.CreatedAt,
		&m.LastUpdated,
	)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

