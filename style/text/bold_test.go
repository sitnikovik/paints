package text

import (
	"testing"

	"github.com/sitnikovik/paints/internal/ansi"
	ansiTextStyle "github.com/sitnikovik/paints/internal/ansi/style/text"
)

func TestBold(t *testing.T) {
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
			want: ansiTextStyle.Bold.String() + "foo" + ansi.Reset.String(),
		},
		{
			name: "empty string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Bold(tt.args.s); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
