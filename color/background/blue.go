package background

import (
	ansi "github.com/sitnikovik/paints/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/ansi/color/background"
)

// Blue returns a new string with blue background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Blue(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Blue + s + ansi.Reset
}
