// Copyright (c) 2025 Ilya Sitnikov
// SPDX-License-Identifier: MIT
package text

import (
	ansi "github.com/sitnikovik/paints/internal/ansi/color"
	ansiText "github.com/sitnikovik/paints/internal/ansi/color/text"
)

// Cyan returns a new string with cyan text.
//
// Wraps the provided string with ANSI codes on the left and right.
func Cyan(s string) string {
	if s == "" {
		return ""
	}
	return ansiText.Cyan + s + ansi.Reset
}
