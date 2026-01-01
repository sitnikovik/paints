// Copyright (c) 2025 Ilya Sitnikov
// SPDX-License-Identifier: MIT
package background

import (
	ansi "github.com/sitnikovik/paints/internal/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Blue returns a new string with blue background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Blue(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Blue + s + ansi.Reset
}
