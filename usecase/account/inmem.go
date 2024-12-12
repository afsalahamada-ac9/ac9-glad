/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package account

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	"strings"
)

// inmem in memory repo
type inmem struct {
	m map[id.ID]*entity.Account
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[id.ID]*entity.Account{}
	return &inmem{
		m: m,
	}
}

// Create an account
func (r *inmem) Create(e *entity.Account) error {
	r.m[e.ID] = e
	return nil
}

// Get retrieves an account
func (r *inmem) Get(id id.ID) (*entity.Account, error) {
	for _, j := range r.m {
		if j.ID == id {
			return r.m[j.ID], nil
		}
	}

	return nil, glad.ErrNotFound
}

// GetByName retrieves an account using username
func (r *inmem) GetByName(tenantID id.ID, username string) (*entity.Account, error) {
	for _, j := range r.m {
		if j.Username == username && j.TenantID == tenantID {
			return r.m[j.ID], nil
		}
	}

	return nil, glad.ErrNotFound
}

// Update an account
func (r *inmem) Update(e *entity.Account) error {
	account := r.m[e.ID]
	if account == nil {
		return glad.ErrNotFound
	}

	account.ExtID = e.ExtID
	account.Username = e.Username
	account.Type = e.Type
	account.CreatedAt = e.CreatedAt
	account.UpdatedAt = e.UpdatedAt

	r.m[e.ID] = account
	return nil
}

// List list accounts
func (r *inmem) List(tenantID id.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	var d []*entity.Account
	for _, j := range r.m {
		// TenantID check removed
		d = append(d, j)
	}
	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit
		if start > len(d) {
			return []*entity.Account{}, nil
		}
		if end > len(d) {
			end = len(d)
		}
		return d[start:end], nil
	}
	return d, nil
}

// Delete deletes an account
func (r *inmem) Delete(id id.ID) error {
	account, err := r.Get(id)
	if err != nil {
		return err
	}

	r.m[account.ID] = nil
	delete(r.m, account.ID)
	return nil
}

// DeleteByName deletes an account using username
func (r *inmem) DeleteByName(tenantID id.ID, username string) error {
	account, err := r.GetByName(tenantID, username)
	if err != nil {
		return err
	}

	r.m[account.ID] = nil
	delete(r.m, account.ID)
	return nil
}

// GetCount gets total accounts for a given tenant
func (r *inmem) GetCount(tenantID id.ID) (int, error) {
	count := 0
	for _, j := range r.m {
		if j.TenantID == tenantID {
			count++
		}
	}
	return count, nil
}

// Search search accounts
func (r *inmem) Search(tenantID id.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	var d []*entity.Account
	for _, j := range r.m {
		if j.TenantID == tenantID &&
			strings.Contains(strings.ToLower(j.Username), query) {
			d = append(d, j)
		}
	}

	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit
		if start > len(d) {
			return []*entity.Account{}, nil
		}
		if end > len(d) {
			end = len(d)
		}
		return d[start:end], nil
	}
	return d, nil
}

// GetByEmail retrieves an account using email
func (r *inmem) GetByEmail(tenantID id.ID, email string) (*entity.Account, error) {
	for _, j := range r.m {
		if j.Email == email && j.TenantID == tenantID {
			return r.m[j.ID], nil
		}
	}

	return nil, glad.ErrNotFound
}
