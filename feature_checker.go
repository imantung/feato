package typfeat

// FeatureChecker resposible to decide if toggle on or off
type FeatureChecker struct {
	listed map[string]*Feature
}

// NewFeatureChecker create new instance of FeatureRouter
func NewFeatureChecker() *FeatureChecker {
	return &FeatureChecker{
		listed: make(map[string]*Feature),
	}
}

// IsEnabled decide is feature enabled
func (r *FeatureChecker) IsEnabled(feature Feature) bool {
	listedFeature, ok := r.listed[feature.Name]
	return (!ok && feature.DefaultState) || (ok && listedFeature.State)
}

// Put Set state for feature
func (r *FeatureChecker) Put(name string, state bool) error {
	listedFeature, ok := r.listed[name]
	if !ok {
		r.listed[name] = &Feature{
			Name:  name,
			State: state,
		}
	} else {
		listedFeature.State = state
	}
	return nil
}

// Get feature by name
func (r *FeatureChecker) Get(name string) *Feature {
	return r.listed[name]
}

// Delete feature by name
func (r *FeatureChecker) Delete(name string) {
	delete(r.listed, name)
}

// Reset the checker
func (r *FeatureChecker) Reset() {
	r.listed = make(map[string]*Feature)
}
