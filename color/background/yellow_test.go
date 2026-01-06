package background

import (
	"testing"

	"github.com/sitnikovik/paints/internal/ansi"
	bg "github.com/sitnikovik/paints/internal/ansi/color/background"
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
			want: bg.Yellow.String() + "foo" + ansi.Reset.String(),
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
