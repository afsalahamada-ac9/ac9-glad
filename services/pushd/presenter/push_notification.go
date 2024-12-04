/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

type PushNotifyInfo struct {
	PushToken    string                 `json:"pushToken"`
	RevokeID     string                 `json:"revokeID"`
	AppVersion   string                 `json:"appVersion"`
	DeviceInfo   map[string]interface{} `json:"deviceInfo"`
	PlatformInfo map[string]interface{} `json:"platformInfo"`
}
