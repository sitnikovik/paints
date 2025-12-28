package text

import (
	ansi "github.com/sitnikovik/paints/ansi/color"
	ansiText "github.com/sitnikovik/paints/ansi/color/text"
)

// Red returns a new string with red text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Red(s string) string {
	if s == "" {
		return ""
	}
	return ansiText.Red + s + ansi.Reset
}
