// Copyright (c) 2025 Ilya Sitnikov
// SPDX-License-Identifier: MIT
package background

import (
	ansi "github.com/sitnikovik/paints/internal/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

// Green returns a new string with green background.
//
// Wraps the provided string with ANSI codes on the left and right.
func Green(s string) string {
	if s == "" {
		return ""
	}
	return ansiBackground.Green + s + ansi.Reset
}
