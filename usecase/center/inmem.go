/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package center

import (
	"strings"

	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
)

// inmem in memory repo
type inmem struct {
	m map[id.ID]*entity.Center
}

// newInmem create new repository
func newInmem() *inmem {
	var m = map[id.ID]*entity.Center{}
	return &inmem{
		m: m,
	}
}

// Create a center
func (r *inmem) Create(e *entity.Center) (id.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

// Get a center
func (r *inmem) Get(id id.ID) (*entity.Center, error) {
	if r.m[id] == nil {
		return nil, glad.ErrNotFound
	}
	return r.m[id], nil
}

// Update a center
func (r *inmem) Update(e *entity.Center) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

// Search centers
func (r *inmem) Search(tenantID id.ID,
	query string,
	page, limit int,
) ([]*entity.Center, error) {
	var centers []*entity.Center
	for _, j := range r.m {
		if j.TenantID == tenantID &&
			strings.Contains(strings.ToLower(j.Name), query) {
			centers = append(centers, j)
		}
	}

	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit

		if start > len(centers) {
			return []*entity.Center{}, nil
		}
		if end > len(centers) {
			end = len(centers)
		}
		return centers[start:end], nil
	}

	return centers, nil
}

// List centers
func (r *inmem) List(tenantID id.ID, page, limit int) ([]*entity.Center, error) {
	var centers []*entity.Center
	for _, j := range r.m {
		if j.TenantID == tenantID {
			centers = append(centers, j)
		}
	}

	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit
		if start > len(centers) {
			return []*entity.Center{}, nil
		}

		if end > len(centers) {
			end = len(centers)
		}
		return centers[start:end], nil
	}
	return centers, nil
}

// Delete a center
func (r *inmem) Delete(id id.ID) error {
	if r.m[id] == nil {
		return glad.ErrNotFound
	}
	r.m[id] = nil
	delete(r.m, id)
	return nil
}

// GetCount gets total centers for a given tenant
func (r *inmem) GetCount(tenantID id.ID) (int, error) {
	count := 0
	for _, j := range r.m {
		if j.TenantID == tenantID {
			count++
		}
	}
	return count, nil
}
