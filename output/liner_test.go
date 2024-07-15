package output

import (
	"reflect"
	"testing"
)

func TestFndLine(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want []int
	}{
		{"%", '%', []int{47, 48, 49, 50, 51, 52, 53, 54}},         // Adjusted expected values based on the old logic
		{"V", 'V', []int{488, 489, 490, 491, 492, 493, 494, 495}}, // Adjusted expected values based on the old logic
		// Add more test cases as needed for other runes or characters
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FndLine(tt.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FndLine(%q) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}

func TestProcessLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		{"Simple Case", "ABC", []int{299, 300, 301, 302, 303, 304, 305, 306, 308, 309, 310, 311, 312, 313, 314, 315, 317, 318, 319, 320, 321, 322, 323, 324}}, // Adjusted expected values based on the old logic
		// Add more test cases as needed for other input strings
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessLine(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessLine(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
