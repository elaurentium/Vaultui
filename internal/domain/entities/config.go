package entities

type Config struct {
	GlobalSettings map[string]interface{} `json:"global_settings"`
	FeatureFlags   map[string]bool        `json:"feature_flags"`
}
