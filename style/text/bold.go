package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	style "github.com/sitnikovik/paints/internal/ansi/style/text"
)

// Bold returns a new string with bold text style.
//
// Wraps the provided string with ANSI codes on the left and right.
func Bold(s string) string {
	if s == "" {
		return ""
	}
	return ansi.WrapString(s, style.Bold)
}
