/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"ac9/glad/services/ldsd/entity"
	"database/sql"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// LiveDarshanPGSQL pgsql repo
type LiveDarshanPGSQL struct {
	db *sql.DB
}

// NewLiveDarshanPGSQL create new repository
func NewLiveDarshanPGSQL(db *sql.DB) *LiveDarshanPGSQL {
	return &LiveDarshanPGSQL{
		db: db,
	}
}

// Create creates a live darshan event
func (r *LiveDarshanPGSQL) Create(ld *entity.LiveDarshan) error {
	query := `
		INSERT INTO live_darshan
			(id, tenant_id, date, start_time, meeting_url, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		l.Log.Errorf("err=%#v", err)
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		ld.ID,
		ld.TenantID,
		ld.Date,
		ld.StartTime,
		ld.MeetingURL,
		ld.CreatedBy,
	)
	if err != nil {
		l.Log.Errorf("err=%#v", err)
	}

	return err
}

// Get retrieves live darshan event using id
func (r *LiveDarshanPGSQL) Get(ldID int64) (*entity.LiveDarshan, error) {
	query := `
		SELECT id, tenant_id, date, start_time, meeting_url, created_by
		FROM live_darshan
		WHERE id = $1
	`

	ld := &entity.LiveDarshan{}
	err := r.db.QueryRow(query, ldID).Scan(
		&ld.ID,
		&ld.TenantID,
		&ld.Date,
		&ld.StartTime,
		&ld.MeetingURL,
		&ld.CreatedBy,
	)
	if err != nil {
		l.Log.Errorf("err=%#v", err)
		return nil, err
	}

	return ld, nil
}

// List lists all live darshan events
func (r *LiveDarshanPGSQL) List(
	tenantID id.ID,
	page, limit int,
) ([]*entity.LiveDarshan, error) {
	query := `
		SELECT id, tenant_id, date, start_time, meeting_url, created_by
		FROM live_darshan
		WHERE tenant_id = $1
	`

	// Add pagination if specified
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query += ` LIMIT $2 OFFSET $3;`
		stmt, err := r.db.Prepare(query)
		if err != nil {
			l.Log.Errorf("err=%#v", err)
			return nil, err
		}
		rows, err := stmt.Query(tenantID, limit, offset)
		if err != nil {
			l.Log.Errorf("err=%#v", err)
			return nil, err
		}
		defer rows.Close()
		return r.scanRows(rows)
	}

	stmt, err := r.db.Prepare(query + ";")
	if err != nil {
		l.Log.Errorf("err=%#v", err)
		return nil, err
	}

	rows, err := stmt.Query(tenantID)
	if err != nil {
		l.Log.Errorf("err=%#v", err)
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// Delete deletes a live darshan event
func (r *LiveDarshanPGSQL) Delete(ldID int64) error {
	res, err := r.db.Exec(`DELETE FROM live_darshan WHERE id = $1;`, ldID)
	if err != nil {
		l.Log.Errorf("err=%#v", err)
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total live darshan events
func (r *LiveDarshanPGSQL) GetCount(tenantID id.ID) (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM live_darshan WHERE tenant_id = $1;`)
	if err != nil {
		return 0, err
	}

	var count int
	err = stmt.QueryRow(tenantID).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// scanRows is a helper function to scan rows into live darshan slices
func (r *LiveDarshanPGSQL) scanRows(rows *sql.Rows) ([]*entity.LiveDarshan, error) {
	var lds []*entity.LiveDarshan

	for rows.Next() {
		ld := &entity.LiveDarshan{}
		err := rows.Scan(
			&ld.ID,
			&ld.TenantID,
			&ld.Date,
			&ld.StartTime,
			&ld.MeetingURL,
			&ld.CreatedBy,
		)
		if err != nil {
			l.Log.Errorf("err=%#v", err)
			return nil, err
		}
		lds = append(lds, ld)
	}

	return lds, nil
}
