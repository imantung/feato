package feato

import "sync"

// SimpleToggleRouter Simple toggle router based on feature name
type SimpleToggleRouter struct {
	featureMap map[string]IndexToggle
	mutex      *sync.Mutex
}

// NewSimpleToggleRouter return new instance of SimpleToggleRouter
func NewSimpleToggleRouter() *SimpleToggleRouter {
	return &SimpleToggleRouter{
		featureMap: make(map[string]IndexToggle),
		mutex:      &sync.Mutex{},
	}
}

// Route route feature to expected toggle
func (r *SimpleToggleRouter) Route(feature *Feature) IndexToggle {
	state, ok := r.featureMap[feature.Name]
	if !ok {
		return feature.DefaultIndexToggle
	}

	return state
}

// SetIndexToggle sets state of feature with corresponding name
func (r *SimpleToggleRouter) SetIndexToggle(name string, index IndexToggle) *SimpleToggleRouter {
	r.mutex.Lock()
	r.featureMap[name] = index
	r.mutex.Unlock()

	return r
}

// SetFeature set feature index toggle
func (r *SimpleToggleRouter) SetFeature(feature *Feature, index IndexToggle) *SimpleToggleRouter {
	return r.SetIndexToggle(feature.Name, index)
}

// Enable enable feature
func (r *SimpleToggleRouter) Enable(feature *Feature) *SimpleToggleRouter {
	return r.SetFeature(feature, EnableIndexToggle)
}

// Disable disable feature
func (r *SimpleToggleRouter) Disable(feature *Feature) *SimpleToggleRouter {
	return r.SetFeature(feature, DisableIndexToggle)
}
