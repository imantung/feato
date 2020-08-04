package feato

import "fmt"

type (
	// Flag for toggle
	Flag *bool
	// FlagStore Simple toggle router based on feature name
	FlagStore map[string]Flag
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

//
// FeatureStore
//

// Register the features to store
func (s FlagStore) Register(features []*Feature) {
	for _, feat := range features {
		s.Put(feat.Name, feat.Flag)
		s.registerChild(feat.Name, feat.Childs)
	}
}

func (s FlagStore) registerChild(parent string, childs []*Feature) {
	for _, child := range childs {
		key := fmt.Sprintf("%s.%s", parent, child.Name)
		s.Put(key, child.Flag)
		s.registerChild(key, child.Childs)
	}
}

// Put flag to key index if flag not empty
func (s FlagStore) Put(key string, flag Flag) {
	if flag != nil {
		s[key] = flag
	}
}
