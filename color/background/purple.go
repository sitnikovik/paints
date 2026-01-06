package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	bg "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Purple returns a new string with purple background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Purple(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, bg.Purple)
}
