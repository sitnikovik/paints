package text

import (
	ansi "github.com/sitnikovik/paints/ansi/color"
	ansiText "github.com/sitnikovik/paints/ansi/color/text"
)

// Green returns a new string with green text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Green(s string) string {
	if s == "" {
		return ""
	}
	return ansiText.Green + s + ansi.Reset
}
