/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package product

import (
	"strings"
	"sync"

	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
)

// inmem in memory repo
type inmem struct {
	m   map[id.ID]*entity.Product
	mut *sync.RWMutex
}

// NewInmem creates a new in memory product repository
func NewInmem() *inmem {
	return &inmem{
		m:   make(map[id.ID]*entity.Product),
		mut: &sync.RWMutex{},
	}
}

// Create stores a product in memory
func (r *inmem) Create(e *entity.Product) (id.ID, error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	r.m[e.ID] = e
	return e.ID, nil
}

// Get retrieves a product from memory
func (r *inmem) Get(id id.ID) (*entity.Product, error) {
	r.mut.RLock()
	defer r.mut.RUnlock()

	if product, ok := r.m[id]; ok {
		return product, nil
	}
	return nil, glad.ErrNotFound
}

// Update updates a product in memory
func (r *inmem) Update(e *entity.Product) error {
	r.mut.Lock()
	defer r.mut.Unlock()

	_, ok := r.m[e.ID]
	if !ok {
		return glad.ErrNotFound
	}

	r.m[e.ID] = e
	return nil
}

// List returns all products from memory for the specified tenant
func (r *inmem) List(tenantID id.ID, page, limit int) ([]*entity.Product, error) {
	r.mut.RLock()
	defer r.mut.RUnlock()

	var products []*entity.Product
	for _, product := range r.m {
		if product.TenantID == tenantID {
			products = append(products, product)
		}
	}

	// Handle pagination if needed
	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit
		if start > len(products) {
			return []*entity.Product{}, nil
		}
		if end > len(products) {
			end = len(products)
		}
		return products[start:end], nil
	}

	return products, nil
}

// Delete marks a product as deleted in memory
func (r *inmem) Delete(id id.ID) error {
	r.mut.Lock()
	defer r.mut.Unlock()

	if _, ok := r.m[id]; ok {
		r.m[id] = nil
		delete(r.m, id)
		return nil
	}
	return glad.ErrNotFound
}

// Search searches for products in memory
func (r *inmem) Search(tenantID id.ID, query string, page, limit int) ([]*entity.Product, error) {
	r.mut.RLock()
	defer r.mut.RUnlock()

	var products []*entity.Product
	for _, product := range r.m {
		if product.TenantID == tenantID &&
			(strings.Contains(strings.ToLower(product.ExtName), strings.ToLower(query)) ||
				strings.Contains(strings.ToLower(product.Title), strings.ToLower(query))) {
			products = append(products, product)
		}
	}

	// Handle pagination if needed
	if page > 0 && limit > 0 {
		start := (page - 1) * limit
		end := start + limit
		if start > len(products) {
			return []*entity.Product{}, nil
		}
		if end > len(products) {
			end = len(products)
		}
		return products[start:end], nil
	}

	return products, nil
}

// GetCount returns count of products for a specific tenant
func (r *inmem) GetCount(tenantID id.ID) (int, error) {
	r.mut.RLock()
	defer r.mut.RUnlock()

	count := 0
	for _, product := range r.m {
		if product.TenantID == tenantID {
			count++
		}
	}
	return count, nil
}

// Upsert upserts a product in memory
func (r *inmem) Upsert(e *entity.Product) (id.ID, error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	for _, product := range r.m {
		if product.ExtID == e.ExtID {
			e.ID = product.ID
		}
	}
	r.m[e.ID] = e
	return e.ID, nil
}

// GetByExtID retrieves id using external id
func (r *inmem) GetByExtID(tenantID id.ID, extID string) (*entity.Product, error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	for _, product := range r.m {
		if product.TenantID == tenantID && product.ExtID == extID {
			return product, nil
		}
	}
	return nil, glad.ErrNotFound
}

// Additional helper methods for testing
func (r *inmem) Clean() {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.m = make(map[id.ID]*entity.Product)
}

func (r *inmem) Count() int {
	r.mut.RLock()
	defer r.mut.RUnlock()
	return len(r.m)
}
