package text

import (
	"testing"

	"github.com/sitnikovik/paints/internal/ansi"
	style "github.com/sitnikovik/paints/internal/ansi/style/text"
)

func TestUnderline(t *testing.T) {
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
			want: style.Underline.String() + "foo" + ansi.Reset.String(),
		},
		{
			name: "empty string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Underline(tt.args.s); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
