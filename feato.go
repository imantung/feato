package feato

// Instance instance of feature manager
var Instance Manager

// SetToggleRouter set toggle router
func SetToggleRouter(toggleRouter ToggleRouter) {
	Instance.ToggleRouter = toggleRouter
}

// Run function according feature
func Run(feature *Feature, runFuncs ...RunFunc) error {
	return Instance.Run(feature, runFuncs...)
}
