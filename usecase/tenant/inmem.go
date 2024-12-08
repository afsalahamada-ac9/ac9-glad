/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package tenant

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
)

// inmem in memory repo
type inmem struct {
	m map[id.ID]*entity.Tenant
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[id.ID]*entity.Tenant{}
	return &inmem{
		m: m,
	}
}

// Create a tenant
func (r *inmem) Create(e *entity.Tenant) (id.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get a tenant
func (r *inmem) Get(id id.ID) (*entity.Tenant, error) {
	if r.m[id] == nil {
		return nil, glad.ErrNotFound
	}
	return r.m[id], nil
}

// Get a tenant by name
func (r *inmem) GetByName(name string) (*entity.Tenant, error) {
	for _, j := range r.m {
		if j.Name == name {
			return j, nil
		}
	}
	return nil, glad.ErrNotFound
}

// Update a tenant
func (r *inmem) Update(e *entity.Tenant) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// List tenants
func (r *inmem) List(page, limit int) ([]*entity.Tenant, error) {
	var tenants []*entity.Tenant
	for _, j := range r.m {
		tenants = append(tenants, j)
	}
	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit
		if start > len(tenants) {
			return []*entity.Tenant{}, nil
		}
		if end > len(tenants) {
			end = len(tenants)
		}

		return tenants[start:end], nil
	}

	return tenants, nil
}

// Delete a tenant
func (r *inmem) Delete(id id.ID) error {
	if r.m[id] == nil {
		return glad.ErrNotFound
	}
	r.m[id] = nil
	delete(r.m, id)
	return nil
}

// GetCount gets total tenants
func (r *inmem) GetCount() (int, error) {
	return len(r.m), nil
}
