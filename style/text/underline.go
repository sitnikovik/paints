package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	style "github.com/sitnikovik/paints/internal/ansi/style/text"
)

// Underline returns a new string with underline text style.
//
// Wraps the provided string with ANSI codes on the left and right.
func Underline(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, style.Underline)
}
