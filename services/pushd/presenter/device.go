/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/pkg/id"
	"ac9/glad/pkg/util"
	"ac9/glad/services/pushd/entity"
	"encoding/json"

	l "ac9/glad/pkg/logger"

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

type Device struct {
	ID         id.ID   `json:"id"`
	AccountID  id.ID   `json:"accountID"`
	TenantID   id.ID   `json:"tenantID"`
	PushToken  string  `json:"pushToken"`
	RevokeID   *string `json:"revokeID,omitempty"`
	AppVersion string  `json:"appVersion"`
}

type NotificationMessage struct {
	Header  string `json:"header"`
	Content string `json:"content"`
}
type Notify struct {
	NotificationMessage
	TenantID  id.ID   `json:"tenantID"`
	AccountID []id.ID `json:"accountID"`
}

type NotifyStatus struct {
	AccountID id.ID `json:"accountID"`
	Status    int   `json:"status"`
}

// ToDevice creates device entity from device register request
func (drr DeviceRegisterRequest) ToDevice(tenantID id.ID, accountID id.ID) (entity.Device, error) {

	var device entity.Device
	deepcopier.Copy(drr).To(&device)

	device.TenantID = tenantID
	device.AccountID = accountID

	deviceInfo, err := json.Marshal(drr.DeviceInfo)
	if err != nil {
		l.Log.Warnf("Unable to marshal DeviceInfo=%#v", drr.DeviceInfo)
		return device, err
	}
	platformInfo, err := json.Marshal(drr.PlatformInfo)
	if err != nil {
		l.Log.Warnf("Unable to marshal PlatformInfo=%#v", drr.PlatformInfo)
		return device, err
	}

	device.DeviceInfo = util.NewString(string(deviceInfo))
	device.PlatformInfo = util.NewString(string(platformInfo))

	l.Log.Warnf("drr=%#v, device=%#v", drr, device)
	return device, nil
}
