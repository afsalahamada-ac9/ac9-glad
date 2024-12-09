/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/pkg/id"
	"ac9/glad/services/pushd/entity"

	"github.com/ulule/deepcopier"
)

type DeviceRegisterRequest struct {
	PushToken    string                 `json:"pushToken"`
	RevokeID     string                 `json:"revokeID"`
	AppVersion   string                 `json:"appVersion"`
	DeviceInfo   map[string]interface{} `json:"deviceInfo"`
	PlatformInfo map[string]interface{} `json:"platformInfo"`
}

type DeviceRegisterResponse struct {
	ID id.ID `json:"id"`
}

// ToDevice creates device entity from device register request
func (drr DeviceRegisterRequest) ToDevice(tenantID id.ID, accountID id.ID) (entity.Device, error) {

	var device entity.Device
	deepcopier.Copy(drr).To(&device)

	device.TenantID = tenantID
	device.AccountID = accountID

	return device, nil
}
