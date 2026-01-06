package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	color "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Red returns a new string with red text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Red(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, color.Red)
}
