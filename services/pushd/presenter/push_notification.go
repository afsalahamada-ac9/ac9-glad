package presenter

type PushNotification struct {
	PushToken   string                 `json:"pushToken"`
	RevokeID    string                 `json:"revokeID"`
	AppVersion  string                 `json:"appVersion"`
	DeviceInfo  map[string]interface{} `json:"deviceInfo"`
	PlatformInfo map[string]interface{} `json:"platformInfo"`
}