package feato

import "fmt"

type (
	// FlagStore interface for flag storage
	FlagStore interface {
		Put(key string, flag Flag)
		IsEnabled(name string) bool
		ClearAll()
	}
)

var (
	// Instance for global use
	Instance FlagStore = make(FlagMap, 0)
)

// RegisterGlobal register features to global instance
func RegisterGlobal(features []*Feature) {
	Register(Instance, features)
}

// Register the features to store
func Register(store FlagStore, features []*Feature) {
	for _, feat := range features {
		store.Put(feat.Name, feat.Flag)
		registerChild(store, feat.Name, feat.Childs)
	}
}

func registerChild(store FlagStore, parent string, childs []*Feature) {
	for _, child := range childs {
		key := fmt.Sprintf("%s.%s", parent, child.Name)
		store.Put(key, child.Flag)
		registerChild(store, key, child.Childs)
	}
}

// IsEnabled return flag by name
func IsEnabled(name string) bool {
	return Instance.IsEnabled(name)
}

// Bool convert flag to bool
func Bool(f Flag) bool {
	if f == nil {
		return false
	}
	return *f
}
