package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Purple returns a new string with purple background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Purple(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Purple.String() + s + ansi.Reset.String()
}
