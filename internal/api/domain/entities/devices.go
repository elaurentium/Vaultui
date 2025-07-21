package entities

type Devices struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	PushToken  string `json:"pushToken"`
	DeviceKey  string `json:"deviceKey"`
}

type DeviceKeysRequest struct {
	DeviceKey string `json:"deviceKey"`
}

type DeviceTokenRequest struct {
	Token string `json:"token"`
}

type WebPushAuthRequest struct {
	Endpoint string `json:"endpoint"`
	P256dh   string `json:"p256dh"`
	Auth     string `json:"auth"`
}

type OtherDeviceKeysUpdateRequest struct {
	Identifier string `json:"identifier"`
	DeviceKey  string `json:"deviceKey"`
}

type UpdateDevicesTrustRequest struct {
	Secret        string                         `json:"secret"`
	CurrentDevice string                         `json:"currentDevice"`
	OtherDevices  []OtherDeviceKeysUpdateRequest `json:"otherDevices,omitempty"`
}

type UntrustDevicesRequest struct {
	Devices []string `json:"devices"`
}

type DeviceResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
}

type ProtectedDeviceResponse struct {
	ID        string `json:"id"`
	DeviceKey string `json:"deviceKey"`
}

type DeviceAuthRequestResponse struct {
	ID        string `json:"id"`
	AuthToken string `json:"authToken"`
}

type DeviceAuthListResponse struct {
	Data []DeviceAuthRequestResponse `json:"data"`
}
