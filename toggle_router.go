package feato

import "sync"

const (
	// Enabled index
	Enabled Index = 0
	// Disabled index
	Disabled Index = 1
)

type (
	// Router responsible to route
	Router interface {
		Route(*Feature) Index
	}
	// Index route
	Index int
	// SimpleToggleRouter Simple toggle router based on feature name
	SimpleToggleRouter struct {
		featureMap map[string]Index
		mutex      *sync.Mutex
	}
)

// NewSimpleToggleRouter return new instance of SimpleToggleRouter
func NewSimpleToggleRouter() *SimpleToggleRouter {
	return &SimpleToggleRouter{
		featureMap: make(map[string]Index),
		mutex:      &sync.Mutex{},
	}
}

// Route route feature to expected toggle
func (r *SimpleToggleRouter) Route(feature *Feature) Index {
	state, ok := r.featureMap[feature.ID]
	if !ok {
		return feature.DefaultIndex
	}

	return state
}

// SetIndexToggle sets state of feature with corresponding name
func (r *SimpleToggleRouter) SetIndexToggle(id string, index Index) *SimpleToggleRouter {
	r.mutex.Lock()
	r.featureMap[id] = index
	r.mutex.Unlock()
	return r
}

// SetFeature set feature index toggle
func (r *SimpleToggleRouter) SetFeature(feature *Feature, index Index) *SimpleToggleRouter {
	return r.SetIndexToggle(feature.ID, index)
}

// Enable enable feature
func (r *SimpleToggleRouter) Enable(feature *Feature) *SimpleToggleRouter {
	return r.SetFeature(feature, Enabled)
}

// Disable disable feature
func (r *SimpleToggleRouter) Disable(feature *Feature) *SimpleToggleRouter {
	return r.SetFeature(feature, Disabled)
}
