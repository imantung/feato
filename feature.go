package typfeat

// Feature model
type Feature struct {
	Name         string `json:"name"`
	DefaultState bool   `json:"default_state"`
	State        bool   `json:"state"`
}
