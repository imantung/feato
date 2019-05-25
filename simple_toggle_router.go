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
func (r *SimpleToggleRouter) SetIndexToggle(name string, state IndexToggle) *SimpleToggleRouter {
	r.mutex.Lock()
	r.featureMap[name] = state
	r.mutex.Unlock()

	return r
}

// Enable enable feature by name
func (r *SimpleToggleRouter) Enable(name string) *SimpleToggleRouter {
	return r.SetIndexToggle(name, EnableIndexToggle)
}

// Disable disable feature by name
func (r *SimpleToggleRouter) Disable(name string) *SimpleToggleRouter {
	return r.SetIndexToggle(name, DisableIndexToggle)
}
