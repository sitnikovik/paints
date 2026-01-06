package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Yellow returns a new string with yellow background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Yellow(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Yellow.String() + s + ansi.Reset.String()
}
