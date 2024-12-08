/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package center

import (
	"testing"
	"time"

	"ac9/glad/entity"
	"ac9/glad/pkg/id"

	"github.com/stretchr/testify/assert"
)

const (
	centerDefault id.ID = 13790493495087071234
	tenantAlice   id.ID = 13790492210917015554
	aliceExtID          = "000aliceExtID"

	// todo: add multi-tenant support
	// tenantBob    id.ID = 13790492210917015555
	bobExtID = "000bobExtID"
)

func newFixtureCenter() *entity.Center {
	return &entity.Center{
		ID:        centerDefault,
		TenantID:  tenantAlice,
		ExtID:     aliceExtID,
		ExtName:   "L-0008",
		Name:      "default1",
		Mode:      entity.CenterInPerson,
		IsEnabled: true,
		CreatedAt: time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	tmpl := newFixtureCenter()
	_, err := m.CreateCenter(tmpl.TenantID, tmpl.ExtID, tmpl.ExtName, tmpl.Name, tmpl.Mode, tmpl.IsEnabled)
	assert.Nil(t, err)
	assert.False(t, tmpl.CreatedAt.IsZero())
}

func Test_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	tmpl1 := newFixtureCenter()
	tmpl2 := newFixtureCenter()
	tmpl2.Name = "default2"
	tmpl2.ExtID = bobExtID

	tID, _ := m.CreateCenter(tmpl1.TenantID, tmpl1.ExtID, tmpl1.ExtName, tmpl1.Name, tmpl1.Mode, tmpl1.IsEnabled)
	_, _ = m.CreateCenter(tmpl2.TenantID, tmpl2.ExtID, tmpl2.ExtName, tmpl2.Name, tmpl2.Mode, tmpl2.IsEnabled)

	t.Run("search", func(t *testing.T) {
		res, err := m.SearchCenters(tmpl1.TenantID, "default1", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(res))
		assert.Equal(t, tmpl1.ExtID, res[0].ExtID)
		assert.Equal(t, tmpl1.Mode, res[0].Mode)

		// 'default' query value matches both the center names
		res, err = m.SearchCenters(tmpl1.TenantID, "default", 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(res))

		res, err = m.SearchCenters(tmpl1.TenantID, "non-existent", 0, 0)
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, res)
	})
	t.Run("list all", func(t *testing.T) {
		all, err := m.ListCenters(tmpl1.TenantID, 0, 0)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("get", func(t *testing.T) {
		saved, err := m.GetCenter(tID)
		assert.Nil(t, err)
		assert.Equal(t, tmpl1.TenantID, saved.TenantID)
		assert.Equal(t, tmpl1.ExtID, saved.ExtID)
		assert.Equal(t, tmpl1.Mode, saved.Mode)
		assert.Equal(t, tmpl1.Name, saved.Name)
	})
}

func Test_Update(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	tmpl := newFixtureCenter()
	id, err := m.CreateCenter(tmpl.TenantID, tmpl.ExtID, tmpl.ExtName, tmpl.Name, tmpl.Mode, tmpl.IsEnabled)
	assert.Nil(t, err)

	saved, _ := m.GetCenter(id)
	saved.Mode = entity.CenterOnline
	assert.Nil(t, m.UpdateCenter(saved))

	updated, err := m.GetCenter(id)
	assert.Nil(t, err)
	assert.Equal(t, entity.CenterOnline, updated.Mode)
}

func TestDelete(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)

	tmpl1 := newFixtureCenter()
	tmpl2 := newFixtureCenter()
	tmpl2.ExtID = bobExtID
	t2ID, _ := m.CreateCenter(tmpl2.TenantID, tmpl2.ExtID, tmpl2.ExtName, tmpl2.Name, tmpl2.Mode, tmpl2.IsEnabled)

	err := m.DeleteCenter(tmpl1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteCenter(t2ID)
	assert.Nil(t, err)
	_, err = m.GetCenter(t2ID)
	assert.Equal(t, entity.ErrNotFound, err)
}
