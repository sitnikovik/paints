package text

import "github.com/sitnikovik/paints/internal/ansi"

const (
	// Bold - the ANSI escape code for bold text.
	Bold ansi.Code = "\033[1m"
	// Dim - the ANSI escape code for dim text.
	Dim ansi.Code = "\033[2m"
	// Italic - the ANSI escape code for italic text.
	Italic ansi.Code = "\033[3m"
	// Underline - the ANSI escape code for underline text.
	Underline ansi.Code = "\033[4m"
)
