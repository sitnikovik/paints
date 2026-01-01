// Copyright (c) 2025 Ilya Sitnikov
// SPDX-License-Identifier: MIT
package background

import (
	"testing"

	ansi "github.com/sitnikovik/paints/internal/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/internal/ansi/color/background"
)

func TestYellow(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "ok",
			args: args{
				s: "foo",
			},
			want: ansiBackground.Yellow + "foo" + ansi.Reset,
		},
		{
			name: "empty string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Yellow(tt.args.s); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
