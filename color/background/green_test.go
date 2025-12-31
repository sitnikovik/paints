package background

import (
	"testing"

	ansi "github.com/sitnikovik/paints/ansi/color"
	ansiBackground "github.com/sitnikovik/paints/ansi/color/background"
)

func TestGreen(t *testing.T) {
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
			want: ansiBackground.Green + "foo" + ansi.Reset,
		},
		{
			name: "empty string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Green(tt.args.s); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
