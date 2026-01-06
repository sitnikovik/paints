package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiTextStyle "github.com/sitnikovik/paints/internal/ansi/style/text"
)

// Italic returns a new string with italic text style.
//
// Wraps the provided string with ANSI codes on the left and right.
func Italic(s string) string {
	if s == "" {
		return ""
	}
	return ansiTextStyle.Italic.String() + s + ansi.Reset.String()
}
