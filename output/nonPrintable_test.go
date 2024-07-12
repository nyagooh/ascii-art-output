package output

import "testing"

func TestNonPrintable(t *testing.T) {

	tests := []struct {
		name string
		line string
		want string
	}{
		{"tab",
			"hello\\tworld",
			"hello    world",
		},
		{"backspace",
			"hello\\bworld",
			"hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonPrintable(tt.line); got != tt.want {
				t.Errorf("NonPrintable() = %v, want %v", got, tt.want)
			}
		})
	}
}
