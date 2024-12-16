package product

import (
	"strings"
	"time"

	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
)

// Service product usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateProduct creates a product
func (s *Service) CreateProduct(tenantID id.ID,
	extName string,
	title string,
	ctype string,
	baseProductExtID string,
	durationDays int32,
	visibility entity.ProductVisibility,
	maxAttendees int32,
	format entity.ProductFormat,
	isAutoApprove bool,
) (id.ID, error) {
	p, err := entity.NewProduct(tenantID,
		extName,
		title,
		ctype,
		baseProductExtID,
		durationDays,
		visibility,
		maxAttendees,
		format,
		isAutoApprove,
	)
	if err != nil {
		return id.IDInvalid, err
	}

	return s.repo.Create(p)
}

// GetProduct retrieves a product
func (s *Service) GetProduct(id id.ID) (*entity.Product, error) {
	p, err := s.repo.Get(id)
	if p == nil {
		return nil, glad.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return p, nil
}

// SearchProducts search product
func (s *Service) SearchProducts(tenantID id.ID, q string, page, limit int) ([]*entity.Product, error) {
	products, err := s.repo.Search(tenantID, strings.ToLower(q), page, limit)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, glad.ErrNotFound
	}
	return products, nil
}

// ListProducts list products
func (s *Service) ListProducts(tenantID id.ID, page, limit int) ([]*entity.Product, error) {
	products, err := s.repo.List(tenantID, page, limit)
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, glad.ErrNotFound
	}
	return products, nil
}

// UpdateProduct Update a product
func (s *Service) UpdateProduct(p *entity.Product) error {
	err := p.Validate()
	if err != nil {
		return err
	}
	p.UpdatedAt = time.Now()
	return s.repo.Update(p)
}

// DeleteProduct Delete a product
func (s *Service) DeleteProduct(id id.ID) error {
	p, err := s.GetProduct(id)
	if p == nil {
		return glad.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// GetCount gets total product count
func (s *Service) GetCount(tenantID id.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}

// UpsertProduct upserts a product
func (s *Service) UpsertProduct(p *entity.Product) (id.ID, error) {
	if p.ID == id.IDInvalid {
		// assign id and during update id should not be overwritten
		p.ID = id.New()

	}

	err := p.Validate()
	if err != nil {
		l.Log.Warnf("err=%v", err)
		return id.IDInvalid, err
	}
	return s.repo.Upsert(p)
}

// GetIDByExtID gets product id using external id
func (s *Service) GetIDByExtID(tenantID id.ID, extID string) (id.ID, error) {
	p, err := s.repo.GetByExtID(tenantID, extID)
	if p == nil {
		l.Log.Warnf("tenantID=%v, extID=%v, err=%v", tenantID, extID, err)
		return id.IDInvalid, glad.ErrNotFound
	}
	if err != nil {
		l.Log.Warnf("tenantID=%v, extID=%v, err=%v", tenantID, extID, err)
		return id.IDInvalid, err
	}

	return p.ID, nil
}
