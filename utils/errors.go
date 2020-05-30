package utils

// NewError implementation
func NewError(text string) error {
	return &RecraftError{text}
}

// RecraftError implementation
type RecraftError struct {
	data string
}

//Error implementation
func (e *RecraftError) Error() string {
	return e.data
}
