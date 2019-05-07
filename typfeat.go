package typfeat

// Run Register the feature and run the function according if eligible
func Run(feat Feature, runFunc func() error) (runnng bool, err error) {
	// TODO:
	return
}

// RunFeature Construct and register the feature and run the function according if eligible
func RunFeature(name string, defaultState bool, runFunc func()) (err error) {
	// TODO:
	return
}

// Enable the feature
func Enable(name string) (err error) {
	// TODO:
	return
}

// Disable the feature
func Disable(name string) (err error) {
	// TODO:
	return
}

// ListedFeatures return listed value
func ListedFeatures() (features []Feature) {
	// TODO:
	return
}
