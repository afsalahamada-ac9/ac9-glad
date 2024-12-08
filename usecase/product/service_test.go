package product

import (
	"testing"
	"time"

	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"

	"github.com/stretchr/testify/assert"
)

const (
	productDefault id.ID = 13790493495087071234
	tenantAlice    id.ID = 13790492210917015554
	aliceExtID           = "000aliceExtID"
	bobExtID             = "000bobExtID"
)

func newFixtureProduct() *entity.Product {
	return &entity.Product{
		ID:            productDefault,
		TenantID:      tenantAlice,
		ExtID:         aliceExtID,
		ExtName:       "default1",
		Title:         "Default Product 1",
		CType:         "workshop",
		Format:        entity.ProductFormatInPerson,
		Visibility:    entity.ProductVisibilityUnlisted,
		MaxAttendees:  20,
		IsAutoApprove: true,
		CreatedAt:     time.Now(),
	}
}

func Test_CreateProduct(t *testing.T) {
	repo := NewInmem()
	m := NewService(repo)
	tmpl := newFixtureProduct()
	_, err := m.CreateProduct(
		tmpl.TenantID,
		tmpl.ExtID,
		tmpl.ExtName,
		tmpl.Title,
		tmpl.CType,
		tmpl.BaseProductExtID,
		tmpl.DurationDays,
		tmpl.Visibility,
		tmpl.MaxAttendees,
		tmpl.Format,
		tmpl.IsAutoApprove,
	)
	assert.Nil(t, err)
	assert.False(t, tmpl.CreatedAt.IsZero())
}

// TODO: Add test cases for page and limit
func Test_SearchAndFind(t *testing.T) {
	repo := NewInmem()
	m := NewService(repo)
	tmpl1 := newFixtureProduct()
	tmpl2 := newFixtureProduct()
	tmpl2.ExtName = "default2"
	tmpl2.Title = "Default Product 2"
	tmpl2.ExtID = bobExtID

	_, _ = m.CreateProduct(
		tmpl1.TenantID,
		tmpl1.ExtID,
		tmpl1.ExtName,
		tmpl1.Title,
		tmpl1.CType,
		tmpl1.BaseProductExtID,
		tmpl1.DurationDays,
		tmpl1.Visibility,
		tmpl1.MaxAttendees,
		tmpl1.Format,
		tmpl1.IsAutoApprove,
	)

	tID, _ := m.CreateProduct(
		tmpl2.TenantID,
		tmpl2.ExtID,
		tmpl2.ExtName,
		tmpl2.Title,
		tmpl2.CType,
		tmpl2.BaseProductExtID,
		tmpl2.DurationDays,
		tmpl2.Visibility,
		tmpl2.MaxAttendees,
		tmpl2.Format,
		tmpl2.IsAutoApprove,
	)

	t.Run("search", func(t *testing.T) {
		res, err := m.SearchProducts(tmpl1.TenantID, "default1", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, tmpl1.ExtID, res[0].ExtID)
		assert.Equal(t, tmpl1.Format, res[0].Format)

		// 'default' query value matches both product names
		res, err = m.SearchProducts(tmpl1.TenantID, "default", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(res))

		res, err = m.SearchProducts(tmpl1.TenantID, "non-existent", 0, 0)
		assert.Equal(t, glad.ErrNotFound, err)
		assert.Nil(t, res)
	})

	t.Run("list all", func(t *testing.T) {
		all, err := m.ListProducts(tmpl1.TenantID, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetProduct(tID)
		assert.Nil(t, err)
		assert.Equal(t, tmpl2.TenantID, saved.TenantID)
		assert.Equal(t, tmpl2.ExtID, saved.ExtID)
		assert.Equal(t, tmpl2.Format, saved.Format)
		assert.Equal(t, tmpl2.ExtName, saved.ExtName)
		assert.Equal(t, tmpl2.Title, saved.Title)
	})
}

func Test_UpdateProduct(t *testing.T) {
	repo := NewInmem()
	m := NewService(repo)
	tmpl := newFixtureProduct()
	id, err := m.CreateProduct(
		tmpl.TenantID,
		tmpl.ExtID,
		tmpl.ExtName,
		tmpl.Title,
		tmpl.CType,
		tmpl.BaseProductExtID,
		tmpl.DurationDays,
		tmpl.Visibility,
		tmpl.MaxAttendees,
		tmpl.Format,
		tmpl.IsAutoApprove,
	)
	assert.Nil(t, err)

	saved, _ := m.GetProduct(id)
	saved.Format = entity.ProductFormatOnline
	assert.Nil(t, m.UpdateProduct(saved))

	updated, err := m.GetProduct(id)
	assert.Nil(t, err)
	assert.Equal(t, entity.ProductFormatOnline, updated.Format)
}

func TestDeleteProduct(t *testing.T) {
	repo := NewInmem()
	m := NewService(repo)

	tmpl1 := newFixtureProduct()
	tmpl2 := newFixtureProduct()
	tmpl2.ExtID = bobExtID

	id2, _ := m.CreateProduct(
		tmpl2.TenantID,
		tmpl2.ExtID,
		tmpl2.ExtName,
		tmpl2.Title,
		tmpl2.CType,
		tmpl2.BaseProductExtID,
		tmpl2.DurationDays,
		tmpl2.Visibility,
		tmpl2.MaxAttendees,
		tmpl2.Format,
		tmpl2.IsAutoApprove,
	)

	err := m.DeleteProduct(tmpl1.ID)
	assert.Equal(t, glad.ErrNotFound, err)

	err = m.DeleteProduct(id2)
	assert.Nil(t, err)
	_, err = m.GetProduct(id2)
	assert.Equal(t, glad.ErrNotFound, err)
}
