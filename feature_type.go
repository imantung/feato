package feato

const (
	Realease FeatureType = iota
	ABTesting
	UserTarget
)

type FeatureType int

func (d FeatureType) String() string {
	return [...]string{"Release", "ABTesting", "UserTarget"}[d]
}
