package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	color "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Purple returns a new string with purple text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Purple(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, color.Purple)
}
