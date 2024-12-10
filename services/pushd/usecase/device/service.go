/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package device

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"ac9/glad/services/pushd/entity"
	"context"
	"database/sql"
	"net/http"
	"strings"
)

// Service device usecase
type Service struct {
	repo DeviceRepository
	pns  PushNS
}

// NewService creates new service
func NewService(r DeviceRepository, pns PushNS) *Service {
	return &Service{
		repo: r,
		pns:  pns,
	}
}

// CreateDevice creates a device
func (s *Service) Create(
	device entity.Device,
) (id.ID, error) {
	d, err := device.New()
	if err != nil {
		return id.IDInvalid, err
	}

	deviceID, err := s.repo.Create(d)
	if err != nil {
		return deviceID, err
	}
	return deviceID, err
}

// GetByAccount retrieves devices
func (s *Service) GetByAccount(tenantID id.ID, accountID id.ID) ([]*entity.Device, error) {
	t, err := s.repo.GetByAccount(tenantID, accountID)
	if t == nil {
		return nil, glad.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return t, nil
}

// DeleteDevice deletes a device
// Note: Since delete is cascaded to dependent tables, no need to call those functions explicitly
func (s *Service) Delete(id id.ID) error {
	err := s.repo.Delete(id)
	switch err {
	case sql.ErrNoRows:
		return glad.ErrNotFound
	default:
		return err
	}
}

// GetCount gets total device count
func (s *Service) GetCount(tenantID id.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}

// Notify notifies all devices registered for the given account
func (s *Service) Notify(tenantID id.ID, accountID []id.ID, header, content string) ([]int, error) {
	var statuses []int
	var rmDeviceList []id.ID

	for _, account := range accountID {
		data, err := s.GetByAccount(tenantID, account)
		if err != nil && err != glad.ErrNotFound {
			statuses = append(statuses, http.StatusNotFound)
			l.Log.Warnf("Account=%v not found. err=%v", account, err)
			continue
		}

		if data == nil {
			statuses = append(statuses, http.StatusNotFound)
			l.Log.Warnf("No devices registered for %v", account)
			continue
		}

		status := http.StatusInternalServerError
		for _, d := range data {
			err = s.pns.Send(context.Background(), d.PushToken, header, content)
			if err != nil {
				// in case of 404, delete the token; no need to check for error
				if strings.Contains(err.Error(), "registration-token-not-registered") {
					l.Log.Warnf("Token=%v is not registered or invalid. Removing", d.PushToken)
					rmDeviceList = append(rmDeviceList, d.ID)
					status = min(status, http.StatusNotFound)
				} else if strings.Contains(err.Error(), "401") {
					l.Log.Warnf("Token=%v is not authorized. Removing", d.PushToken)
					rmDeviceList = append(rmDeviceList, d.ID)
					status = min(status, http.StatusUnauthorized)
				}
				l.Log.Warnf("Unable to send notification to %v, err=%v", accountID, err)
				status = min(status, http.StatusInternalServerError)
			} else {
				status = min(status, http.StatusCreated)
			}
		}
		statuses = append(statuses, status)
	}

	l.Log.Warnf("Device ids=%v to be removed", rmDeviceList)

	// Note: We could optimize this by creating a sql statement with all device ids at once
	for _, deviceID := range rmDeviceList {
		s.repo.Delete(deviceID)
	}

	return statuses, nil
}
