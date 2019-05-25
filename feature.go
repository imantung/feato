package feato

import "github.com/rs/xid"

// Feature model for feature
type Feature struct {
	ID                 string      `json:"name"`
	DefaultIndexToggle IndexToggle `json:"default_index_toggle"`
}

// NewUniqueFeature return new instance of Feature with random unique ID
func NewUniqueFeature() *Feature {
	return NewFeature(xid.New().String())
}

// NewFeature return new instance of Feature with specific ID, by default index toggle is disabled
func NewFeature(id string) *Feature {
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

// EnableByDefault set default index toggle to enable
func (f *Feature) EnableByDefault() *Feature {
	return f.SetDefaultIndexToggle(EnableIndexToggle)
}
