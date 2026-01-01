package text

import (
	ansi "github.com/sitnikovik/paints/internal/ansi/color"
	ansiText "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Purple returns a new string with purple text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Purple(s string) string {
	if s == "" {
		return ""
	}
	return ansiText.Purple + s + ansi.Reset
}
