package feato

import "sync"

type NamelyToggleRouter struct {
	featureMap map[string]bool
	mutex      *sync.Mutex
}

// NewNamelyToggleRouter return new instance of NamelyToggleRouter
func NewNamelyToggleRouter() ToggleRouter {
	return &NamelyToggleRouter{
		featureMap: make(map[string]bool),
		mutex:      &sync.Mutex{},
	}
}

// Route route feature to expected toggle
func (r *NamelyToggleRouter) Route(feature Feature) int {
	if r.WhenEqual(feature) {
		return 0
	}

	return 1
}

// WhenEqual compares feature with corresponding name to featureMap
func (r *NamelyToggleRouter) WhenEqual(feature Feature) bool {
	state, ok := r.featureMap[feature.Name]
	if !ok {
		return feature.DefaultState
	}

	return state
}

// SetState sets state of feature with corresponding name
func (r *NamelyToggleRouter) SetState(name string, state bool) {
	r.mutex.Lock()
	r.featureMap[name] = state
	r.mutex.Unlock()
}

// GetState gets the state of feature with corresponding name
func (r *NamelyToggleRouter) GetState(name string) bool {
	return r.featureMap[name]
}
