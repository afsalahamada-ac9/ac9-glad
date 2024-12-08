/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package device

import (
	"testing"
	"time"

	"ac9/glad/pkg/id"

	"ac9/glad/services/pushd/entity"

	"github.com/stretchr/testify/assert"
)

const (
	deviceDefault id.ID = 13790493495087071234
	deviceIDBob   id.ID = 13790493495087071235
	tenantAlice   id.ID = 13790492210917015554

	aliceAccountID = 13790493495087075501

	// todo: add multi-tenant support
	// tenantBob    id.ID = 13790492210917015555
	// bobAccountID = 13790493495087075502
)

func newFixtureDevice() *entity.Device {
	return &entity.Device{
		ID:           deviceDefault,
		TenantID:     tenantAlice,
		AccountID:    aliceAccountID,
		Token:        "AliceToken",
		RevokeID:     nil,
		AppVersion:   "v2024.12.3",
		DeviceInfo:   nil,
		PlatformInfo: nil,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := newInmemDevice()
	s := NewService(repo)
	device := newFixtureDevice()
	_, err := s.CreateDevice(*device)

	assert.Nil(t, err)
	assert.False(t, device.CreatedAt.IsZero())
}

func Test_GetByAccount(t *testing.T) {
	repo := newInmemDevice()
	s := NewService(repo)
	device1 := newFixtureDevice()
	device2 := newFixtureDevice()
	device2.Token = "BobToken"
	device2.ID = deviceIDBob

	_, _ = s.CreateDevice(*device1)
	_, _ = s.CreateDevice(*device2)

	t.Run("list all", func(t *testing.T) {
		all, err := s.GetByAccount(device1.TenantID, device1.AccountID)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})
}

func TestDelete(t *testing.T) {
	repo := newInmemDevice()
	s := NewService(repo)

	_ = newFixtureDevice()
	device2 := newFixtureDevice()
	bID, _ := s.CreateDevice(*device2)

	err := s.DeleteDevice(bID)
	assert.Nil(t, err)
}
