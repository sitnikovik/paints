package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Green returns a new string with green background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Green(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Green.String() + s + ansi.Reset.String()
}
