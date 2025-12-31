package background

import (
	ansi "github.com/sitnikovik/paints/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/ansi/color/background"
)

// Purple returns a new string with purple background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Purple(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Purple + s + ansi.Reset
}
