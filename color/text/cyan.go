package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiText "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Cyan returns a new string with cyan text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Cyan(s string) string {
	if s == "" {
		return ""
	}
	return ansiText.Cyan.String() + s + ansi.Reset.String()
}
