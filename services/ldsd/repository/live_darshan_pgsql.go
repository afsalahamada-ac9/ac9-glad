/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"ac9/glad/services/ldsd/entity"
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type LiveDarshanPGSQL struct {
	db *sql.DB
}

func NewLiveDarshanPGSQL(db *sql.DB) *LiveDarshanPGSQL {
	return &LiveDarshanPGSQL{
		db: db,
	}
}

func (r *LiveDarshanPGSQL) Create(ld *entity.LiveDarshan) error {
	query := `
		INSERT INTO live_darshan (id, date, start_time, meeting_url, created_by)
		VALUES ($1, $2, $3, $4, $5)
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		ld.ID,
		ld.Date,
		ld.StartTime,
		ld.MeetingURL,
		ld.CreatedBy,
	)
	return err
}

func (r *LiveDarshanPGSQL) Get(id string) (*entity.LiveDarshan, error) {
	query := `
		SELECT id, date, start_time, meeting_url, created_by
		FROM live_darshan
		WHERE id = $1
	`

	ld := &entity.LiveDarshan{}
	err := r.db.QueryRow(query, id).Scan(
		&ld.ID,
		&ld.Date,
		&ld.StartTime,
		&ld.MeetingURL,
		&ld.CreatedBy,
	)
	if err != nil {
		return nil, err
	}

	return ld, nil
}
