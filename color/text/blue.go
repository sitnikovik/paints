package text

import (
	ansi "github.com/sitnikovik/paints/internal/ansi/color"
	ansiText "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Blue returns a new string with blue text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Blue(s string) string {
	if s == "" {
		return ""
	}
	return ansiText.Blue + s + ansi.Reset
}
