package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	style "github.com/sitnikovik/paints/internal/ansi/style/text"
)

// Italic returns a new string with italic text style.
//
// Wraps the provided string with ANSI codes on the left and right.
func Italic(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, style.Italic)
}
