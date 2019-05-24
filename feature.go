package feato

// Feature model
type Feature struct {
	Name         string      `json:"name"`
	DefaultState IndexToggle `json:"default_state"`
	Type         FeatureType `json:"type"`
}
