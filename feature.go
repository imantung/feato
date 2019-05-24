package feato

// Feature model for feature
type Feature struct {
	Name               string      `json:"name"`
	DefaultIndexToggle IndexToggle `json:"default_state"`
}

// NewFeature return new instance of Feature
func NewFeature(name string) *Feature {
	return &Feature{
		Name:               name,
		DefaultIndexToggle: DisableIndexToggle,
	}
}

// SetDefaultIndexToggle set default toggle
func (f *Feature) SetDefaultIndexToggle(defaultIndexToggle IndexToggle) *Feature {
	f.DefaultIndexToggle = defaultIndexToggle
	return f
}
