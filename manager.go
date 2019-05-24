package feato

type Manager struct {
	toggleRouter ToggleRouter
}

// Run run functions according feature and toggleRouter
func (m *Manager) Run(feature Feature, runFuncs ...func() error) error {
	index := m.toggleRouter.Route(feature)
	return runFuncs[index]()
}
