package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	bg "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Blue returns a new string with blue background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Blue(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, bg.Blue)
}
