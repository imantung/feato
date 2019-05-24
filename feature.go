package feato

const (
	Realease FeatureType = iota
	ABTesting
	UserTarget
)

// Feature model
type Feature struct {
	Name         string      `json:"name"`
	DefaultState bool        `json:"default_state"`
	State        bool        `json:"state"`
	Type         FeatureType `json:"type"`
}

type FeatureType int

func (d FeatureType) String() string {
	return [...]string{"Release", "ABTesting", "UserTarget"}[d]
}
