package text

import (
	"github.com/sitnikovik/paints/internal/ansi"
	ansiTextStyle "github.com/sitnikovik/paints/internal/ansi/style/text"
)

// Dim returns a new string with dim text style.
//
// Wraps the provided string with ANSI codes on the left and right.
func Dim(s string) string {
	if s == "" {
		return ""
	}
	return ansiTextStyle.Dim.String() + s + ansi.Reset.String()
}
