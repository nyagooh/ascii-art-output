package output

import "testing"

func TestMaps(t *testing.T) {

	tests := []struct {
		name     string
		args     int
		filename string
		want     string
	}{
		{"T", 470, "standard.txt", " _______  "},
		{"X", 509, "standard.txt", "  > <   "},
		{"[", 540, "standard.txt", "      "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maps(tt.args, tt.filename); got != tt.want {
				t.Errorf("Maps() = %v, want %v", got, tt.want)
			}
		})
	}
}
