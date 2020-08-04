package feato

type (
	// Flag for toggle
	Flag *bool
	// FlagMap map of flag
	FlagMap map[string]Flag
	// Feature model for feature
	Feature struct {
		Name   string
		Flag   Flag
		Childs []*Feature
	}
)

var (
	enabled  = true
	disabled = false
	// Enabled flag
	Enabled Flag = &enabled
	// Disabled flag
	Disabled Flag = &disabled
)

var _ FlagStore = (FlagMap)(nil)

// Put flag to key index if flag not empty
func (m FlagMap) Put(key string, flag Flag) {
	if flag != nil {
		m[key] = flag
	}
}

// IsEnabled return key flag
func (m FlagMap) IsEnabled(key string) bool {
	if _, ok := m[key]; !ok {
		return false
	}
	return Bool(m[key])
}

// ClearAll clear all feature
func (m FlagMap) ClearAll() {
	for k := range m {
		delete(m, k)
	}
}
