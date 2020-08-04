package feato

type (
	// Feature model for feature
	Feature struct {
		ID           string `json:"name"`
		DefaultIndex Index  `json:"default"`
	}
)

// NewFeature return new instance of Feature with specific ID, by default index toggle is disabled
func NewFeature(id string) *Feature {
	return &Feature{
		ID:           id,
		DefaultIndex: Disabled,
	}
}

// SetDefaultIndex set default toggle
func (f *Feature) SetDefaultIndex(defaultIndex Index) *Feature {
	f.DefaultIndex = defaultIndex
	return f
}

// EnableByDefault set default index toggle to enable
func (f *Feature) EnableByDefault() *Feature {
	return f.SetDefaultIndex(Enabled)
}
