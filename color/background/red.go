// Copyright (c) 2025 Ilya Sitnikov
// SPDX-License-Identifier: MIT
package background

import (
	ansi "github.com/sitnikovik/paints/internal/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Red returns a new string with red background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Red(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Red + s + ansi.Reset
}
