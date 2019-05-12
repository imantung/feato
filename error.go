package feato

const ()

// Error model
type Error string

func (e Error) Error() string {
	return string(e)
}
