package background

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Cyan returns a new string with cyan background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Cyan(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Cyan.String() + s + ansi.Reset.String()
}
