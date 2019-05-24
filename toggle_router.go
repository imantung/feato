package feato

type ToggleRouter interface {
	Route(Feature) IndexToggle
}
