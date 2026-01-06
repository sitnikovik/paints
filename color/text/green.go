package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiText "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Green returns a new string with green text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Green(s string) string {
	if s == "" {
		return ""
	}
	return ansiText.Green.String() + s + ansi.Reset.String()
}
