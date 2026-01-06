package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	color "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Yellow returns a new string with yellow text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Yellow(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, color.Yellow)
}
