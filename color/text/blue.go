package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	color "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Blue returns a new string with blue text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Blue(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, color.Blue)
}
