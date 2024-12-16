/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"database/sql"
	"encoding/json"
	"time"

	"ac9/glad/entity"
	"ac9/glad/pkg/common"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
)

// CenterPGSQL mysql repo
type CenterPGSQL struct {
	db *sql.DB
}

// NewCenterPGSQL create new repository
func NewCenterPGSQL(db *sql.DB) *CenterPGSQL {
	return &CenterPGSQL{
		db: db,
	}
}

// Create creates a center
func (r *CenterPGSQL) Create(e *entity.Center) (id.ID, error) {
	stmt, err := r.db.Prepare(`
		INSERT INTO center (id, tenant_id, name, address, geo_location,
		 capacity, mode, web_page, is_national, is_enabled, created_at)
		VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`)
	if err != nil {
		return e.ID, err
	}
	_, err = stmt.Exec(
		e.ID,
		e.TenantID,
		e.Name,
		e.Address,     // TODO: to be converted into json
		e.GeoLocation, // TODO: to be converted into json
		e.Capacity,
		e.Mode,
		e.WebPage,
		e.IsNational,
		e.IsEnabled,
		time.Now().Format("2006-01-02"),
	)
	if err != nil {
		return e.ID, err
	}
	err = stmt.Close()
	if err != nil {
		return e.ID, err
	}
	return e.ID, nil
}

// Get retrieves a center
// Not all fields are required for v1
func (r *CenterPGSQL) Get(id id.ID) (*entity.Center, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, name, ext_name, mode, created_at FROM center WHERE id = $1;`)
	if err != nil {
		return nil, err
	}
	var c entity.Center
	var extName sql.NullString
	var name sql.NullString
	var mode sql.NullString
	err = stmt.QueryRow(id).Scan(&c.ID, &c.TenantID, &name, &extName, &mode, &c.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	c.Name = name.String
	c.ExtName = extName.String
	c.Mode = entity.CenterMode(mode.String)

	return &c, nil
}

// Update updates a center
func (r *CenterPGSQL) Update(e *entity.Center) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`
		UPDATE center SET name = $1, mode = $2, updated_at = $3 WHERE id = $4;`,
		e.Name, e.Mode, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

// Search searches centers
func (r *CenterPGSQL) Search(tenantID id.ID,
	q string, page, limit int,
) ([]*entity.Center, error) {
	query := `
		SELECT id, tenant_id, name, ext_name, capacity, mode, created_at
		FROM center
		WHERE is_enabled = TRUE
			AND tenant_id = $1
			AND (LOWER(name) LIKE LOWER($2)
				OR LOWER(ext_name) LIKE LOWER($2)
			)
	`

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query += ` LIMIT $3 OFFSET $4;`
		stmt, err := r.db.Prepare(query)
		if err != nil {
			return nil, err
		}

		rows, err := stmt.Query(tenantID, "%"+q+"%", limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		return r.scanRows(rows)
	}

	stmt, err := r.db.Prepare(query + ";")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(tenantID, "%"+q+"%")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// List lists centers
func (r *CenterPGSQL) List(tenantID id.ID, page, limit int) ([]*entity.Center, error) {
	query := `
		SELECT id, tenant_id, name, ext_name, capacity, mode, created_at
		FROM center
		WHERE is_enabled = TRUE AND tenant_id = $1
	`
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query += ` LIMIT $2 OFFSET $3;`
		stmt, err := r.db.Prepare(query)
		if err != nil {
			return nil, err
		}

		rows, err := stmt.Query(tenantID, limit, offset)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		return r.scanRows(rows)
	}

	stmt, err := r.db.Prepare(query + ";")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(tenantID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRows(rows)
}

// Delete deletes a center
func (r *CenterPGSQL) Delete(id id.ID) error {
	res, err := r.db.Exec(`DELETE FROM center WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total centers
func (r *CenterPGSQL) GetCount(tenantID id.ID) (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM center WHERE tenant_id = $1;`)
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

func (r *CenterPGSQL) scanRows(rows *sql.Rows) ([]*entity.Center, error) {
	var centers []*entity.Center
	for rows.Next() {
		var center entity.Center
		var ext_name, name sql.NullString
		var capacity sql.NullInt32

		err := rows.Scan(
			&center.ID,
			&center.TenantID,
			&name,
			&ext_name,
			&capacity,
			&center.Mode,
			&center.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		center.Name = name.String
		center.ExtName = ext_name.String
		center.Capacity = capacity.Int32

		centers = append(centers, &center)

	}
	return centers, nil
}

// Upsert inserts or updates the center and returns the id
func (r *CenterPGSQL) Upsert(e *entity.Center) (id.ID, error) {
	stmt, err := r.db.Prepare(`
		WITH upsert AS (
			INSERT INTO center (
				id, tenant_id, ext_id, ext_name, name, address,
				geo_location, capacity, mode, web_page,
				is_national, is_enabled, created_at, updated_at
			)
			VALUES
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
			ON CONFLICT (ext_id)
			DO UPDATE
				SET tenant_id = $2, ext_name = $4, name = $5, address = $6, geo_location = $7,
					capacity = $8, mode = $9, web_page = $10,
					is_national = $11,  is_enabled = $12, created_at = $13, updated_at = $14
			WHERE center.updated_at < $14
			RETURNING id
		)
		SELECT id FROM upsert
		UNION ALL
		SELECT id FROM center WHERE ext_id = $3 AND NOT EXISTS (SELECT 1 FROM upsert);
	`)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return id.IDInvalid, err
	}

	jsonAddress, err := json.Marshal(e.Address)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return id.IDInvalid, err
	}
	l.Log.Debugf("Address=%#v jsonAddress=%v", e.Address, jsonAddress)

	jsonGeoLocation, err := json.Marshal(e.GeoLocation)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return id.IDInvalid, err
	}
	l.Log.Debugf("GeoLocation=%#v jsonGeoLocation=%v", e.GeoLocation, jsonGeoLocation)

	var centerID id.ID
	err = stmt.QueryRow(
		e.ID,
		e.TenantID,
		e.ExtID,
		e.ExtName,
		e.Name,
		string(jsonAddress),
		string(jsonGeoLocation),
		e.Capacity,
		string(e.Mode),
		e.WebPage,
		e.IsNational,
		e.IsEnabled,
		e.CreatedAt.Format(common.DBFormatDateTimeMS),
		e.UpdatedAt.Format(common.DBFormatDateTimeMS),
	).Scan(&centerID)
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return id.IDInvalid, err
	}

	return centerID, nil
}
