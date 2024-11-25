/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package account

import (
	"strings"
	"sudhagar/glad/entity"
)

// inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Account
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.Account{}
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
func (r *inmem) Get(id entity.ID) (*entity.Account, error) {
	for _, j := range r.m {
		if j.ID == id {
			return r.m[j.ID], nil
		}
	}

	return nil, entity.ErrNotFound
}

// Get retrieves an account using username
func (r *inmem) GetByName(tenantID entity.ID, username string) (*entity.Account, error) {
	for _, j := range r.m {
		if j.Username == username && j.TenantID == tenantID {
			return r.m[j.ID], nil
		}
	}

	return nil, entity.ErrNotFound
}

// Update an account
func (r *inmem) Update(e *entity.Account) error {
	account := r.m[e.ID]
	if account == nil {
		return entity.ErrNotFound
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
func (r *inmem) List(tenantID entity.ID) ([]*entity.Account, error) {
	var d []*entity.Account
	for _, j := range r.m {
		// TenantID check removed
		d = append(d, j)
	}
	return d, nil
}

// Delete deletes an account
func (r *inmem) Delete(id entity.ID) error {
	account, err := r.Get(id)
	if err != nil {
		return err
	}

	r.m[account.ID] = nil
	delete(r.m, account.ID)
	return nil
}

// DeleteByName deletes an account using username
func (r *inmem) DeleteByName(tenantID entity.ID, username string) error {
	account, err := r.GetByName(tenantID, username)
	if err != nil {
		return err
	}

	r.m[account.ID] = nil
	delete(r.m, account.ID)
	return nil
}

// GetCount gets total accounts for a given tenant
func (r *inmem) GetCount(tenantID entity.ID) (int, error) {
	count := 0
	for _, j := range r.m {
		if j.TenantID == tenantID {
			count++
		}
	}
	return count, nil
}

// Search search accounts
func (r *inmem) Search(tenantID entity.ID, query string) ([]*entity.Account, error) {
	var d []*entity.Account
	for _, j := range r.m {
		if j.TenantID == tenantID &&
			strings.Contains(strings.ToLower(j.Username), query) {
			d = append(d, j)
		}
	}
	return d, nil
}
