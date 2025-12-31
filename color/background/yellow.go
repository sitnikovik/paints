package background

import (
	ansi "github.com/sitnikovik/paints/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/ansi/color/background"
)

// Yellow returns a new string with yellow background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Yellow(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Yellow + s + ansi.Reset
}
