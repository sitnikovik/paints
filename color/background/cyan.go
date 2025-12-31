package background

import (
	ansi "github.com/sitnikovik/paints/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/ansi/color/background"
)

// Cyan returns a new string with cyan background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Cyan(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Cyan + s + ansi.Reset
}
