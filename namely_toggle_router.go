package feato

import "sync"

// NamelyToggleRouter Simple toggle router based on feature name
type NamelyToggleRouter struct {
	featureMap map[string]IndexToggle
	mutex      *sync.Mutex
}

// NewNamelyToggleRouter return new instance of NamelyToggleRouter
func NewNamelyToggleRouter() *NamelyToggleRouter {
	return &NamelyToggleRouter{
		featureMap: make(map[string]IndexToggle),
		mutex:      &sync.Mutex{},
	}
}

// Route route feature to expected toggle
func (r *NamelyToggleRouter) Route(feature *Feature) IndexToggle {
	state, ok := r.featureMap[feature.Name]
	if !ok {
		return feature.DefaultIndexToggle
	}

	return state
}

// SetIndexToggle sets state of feature with corresponding name
func (r *NamelyToggleRouter) SetIndexToggle(name string, state IndexToggle) {
	r.mutex.Lock()
	r.featureMap[name] = state
	r.mutex.Unlock()
}
