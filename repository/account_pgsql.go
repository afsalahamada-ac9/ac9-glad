/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package repository

import (
	"database/sql"
	"time"

	"sudhagar/glad/entity"
)

// AccountPGSQL mysql repo
type AccountPGSQL struct {
	db *sql.DB
}

// NewAccountPGSQL create new repository
func NewAccountPGSQL(db *sql.DB) *AccountPGSQL {
	return &AccountPGSQL{
		db: db,
	}
}

// Create creates an account
func (r *AccountPGSQL) Create(e *entity.Account) error {
	stmt, err := r.db.Prepare(`
		INSERT INTO account (id, tenant_id, ext_id, username, first_name, last_name, phone, email, type, created_at) 
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		e.ID,
		e.TenantID,
		e.ExtID,
		e.Username,
		e.FirstName,
		e.LastName,
		e.Phone,
		e.Email,
		e.Type,
		time.Now().Format("2006-01-02"),
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

// Note: Accounts are global in nature, but for storage purposes they will be assigned to some tenants.
// GetByName retrieves an account using username
func (r *AccountPGSQL) GetByName(tenantID entity.ID, username string) (*entity.Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, ext_id, type, created_at FROM account WHERE tenant_id = $1 AND username = $2;`)
	if err != nil {
		return nil, err
	}
	var t entity.Account
	var acct_type sql.NullString
	err = stmt.QueryRow(tenantID, username).Scan(&t.ID, &t.ExtID, &acct_type, &t.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	t.Username = username
	t.Type = entity.AccountType(acct_type.String)
	return &t, nil
}

// Update updates an account
func (r *AccountPGSQL) Update(e *entity.Account) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`UPDATE account SET username = $1, type = $2, updated_at = $3 WHERE id = $4;`,
		e.Username, e.Type, e.UpdatedAt.Format("2006-01-02"), e.ID)
	if err != nil {
		return err
	}
	return nil
}

// List accounts
func (r *AccountPGSQL) List(tenantID entity.ID, page, limit int) ([]*entity.Account, error) {
	query := `
		SELECT id, ext_id, username, type, created_at FROM account`
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

// Delete deletes an account
func (r *AccountPGSQL) Delete(id entity.ID) error {
	res, err := r.db.Exec(`DELETE FROM account WHERE id = $1;`, id)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// DeleteByName deletes an account using username
func (r *AccountPGSQL) DeleteByName(tenantID entity.ID, username string) error {
	res, err := r.db.Exec(`DELETE FROM account WHERE tenant_id = $1 AND username = $2;`, tenantID, username)
	if err != nil {
		return err
	}

	if cnt, _ := res.RowsAffected(); cnt == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Get total accounts
func (r *AccountPGSQL) GetCount(tenantID entity.ID) (int, error) {
	stmt, err := r.db.Prepare(`SELECT count(*) FROM account WHERE tenant_id = $1;`)
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

// Get retrieves an account
func (r *AccountPGSQL) Get(id entity.ID) (*entity.Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT id, tenant_id, ext_id, username, first_name, last_name, 
			phone, email, type, created_at
		FROM account WHERE id = $1;`)
	if err != nil {
		return nil, err
	}

	var a entity.Account
	var ext_id, first_name, last_name, phone, email, accountType sql.NullString

	err = stmt.QueryRow(id).Scan(
		&a.ID,
		&a.TenantID,
		&ext_id,
		&a.Username,
		&first_name,
		&last_name,
		&phone,
		&email,
		&accountType,
		&a.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	a.ExtID = ext_id.String
	a.FirstName = first_name.String
	a.LastName = last_name.String
	a.Phone = phone.String
	a.Email = email.String
	a.Type = entity.AccountType(accountType.String)

	return &a, nil
}

// Search searches accounts
func (r *AccountPGSQL) Search(tenantID entity.ID, q string, page, limit int) ([]*entity.Account, error) {
	// OR LOWER(first_name) LIKE LOWER($2)
	// OR LOWER(last_name) LIKE LOWER($2)
	// OR LOWER(email) LIKE LOWER($2)

	query := `
	SELECT id, tenant_id, ext_id, username, first_name, last_name,
	phone, email, type, created_at
	FROM account 
	WHERE tenant_id = $1 
	AND (
		LOWER(username) LIKE LOWER($2) 
		)`

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

func (r *AccountPGSQL) scanRows(rows *sql.Rows) ([]*entity.Account, error) {
	var accounts []*entity.Account
	for rows.Next() {
		var account entity.Account
		var ext_id, username, first_name, last_name, phone, email sql.NullString

		err := rows.Scan(
			&account.ID,
			&account.TenantID,
			&ext_id,
			&username,
			&first_name,
			&last_name,
			&phone,
			&email,
		)

		if err != nil {
			return nil, err
		}
		account.ExtID = ext_id.String
		account.Username = username.String
		account.FirstName = first_name.String
		account.LastName = last_name.String
		account.Phone = phone.String
		account.Email = email.String

		accounts = append(accounts, &account)
	}
	return accounts, nil
}
