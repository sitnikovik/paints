package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Blue returns a new string with blue background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Blue(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Blue.String() + s + ansi.Reset.String()
}
