package feato

const (
	ErrOutOfRunFunctionsIndex = Error("Out of run functions index")
)

// Error model
type Error string

func (e Error) Error() string {
	return string(e)
}
