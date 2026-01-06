package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiText "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Blue returns a new string with blue text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Blue(s string) string {
	if s == "" {
		return ""
	}
	return ansiText.Blue.String() + s + ansi.Reset.String()
}
