package feato

import "github.com/rs/xid"

// Feature model for feature
type Feature struct {
	ID                 string      `json:"name"`
	DefaultIndexToggle IndexToggle `json:"default_index_toggle"`
}

// NewFeature return new instance of Feature with random unique ID
func NewFeature() *Feature {
	return NewFeatureWithID(xid.New().String())
}

// NewFeatureWithID return new instance of Feature with specific ID
func NewFeatureWithID(id string) *Feature {
	return &Feature{
		ID:                 id,
		DefaultIndexToggle: DisableIndexToggle,
	}
}

// SetDefaultIndexToggle set default toggle
func (f *Feature) SetDefaultIndexToggle(defaultIndexToggle IndexToggle) *Feature {
	f.DefaultIndexToggle = defaultIndexToggle
	return f
}
