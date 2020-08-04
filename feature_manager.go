package feato

type FeatureManager struct {
	ToggleRouter Router
}

type RunFunc func() error

// Run run functions according feature and toggleRouter
func (m *FeatureManager) Run(feature *Feature, runFuncs ...RunFunc) (ok bool, err error) {
	index := m.ToggleRouter.Route(feature)
	if len(runFuncs) <= int(index) {
		return
	}

	ok = true
	err = runFuncs[index]()
	return
}
