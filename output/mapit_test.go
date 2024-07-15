package output

import "testing"

func TestMaps(t *testing.T) {
	tests := []struct {
		name string
		args int
		want string
	}{
		{"T", 470, " _______  "},
		{"X", 509, "  > <   "},
		{"[", 540, "      "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maps(tt.args, "standard.txt"); got != tt.want {
				t.Errorf("Maps() = %v, want %v", got, tt.want)
			}
		})
	}
}