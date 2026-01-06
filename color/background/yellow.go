package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	bg "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Yellow returns a new string with yellow background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Yellow(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, bg.Yellow)
}
