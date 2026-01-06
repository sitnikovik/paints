package ansi

const (
	// Reset - the ANSI escape code to reset text formatting.
	Reset Code = "\033[0m"
)

// Code represents an ANSI escape code.
type Code string

// String returns the string representation of the ANSI code.
func (c Code) String() string {
	return string(c)
}
