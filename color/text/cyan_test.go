package text

import (
	"testing"

	ansi "github.com/sitnikovik/paints/internal/ansi/color"
	ansiText "github.com/sitnikovik/paints/internal/ansi/color/text"
)

func TestCyan(t *testing.T) {
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
			want: ansiText.Cyan + "foo" + ansi.Reset,
		},
		{
			name: "empty string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Cyan(tt.args.s); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
