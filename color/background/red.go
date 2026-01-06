package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	bg "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Red returns a new string with red background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Red(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, bg.Red)
}
