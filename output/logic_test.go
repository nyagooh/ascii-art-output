package output

import (
	"os"
	"testing"
)

var tests = []struct {
	name   string
	output string
	input  string
	args   []string
	want   []byte
}{
	{
		"Test1",
		"sample.txt",
		"h",
		[]string{"h", "standard"},
		[]byte{32, 95, 32, 32, 32, 32, 32, 32, 10, 124, 32, 124, 32, 32, 32, 32, 32, 10, 124, 32, 124, 95, 95, 32, 32, 32, 10, 124, 32, 32, 95, 32, 92, 32, 32, 10, 124, 32, 124, 32, 124, 32, 124, 32, 10, 124, 95, 124, 32, 124, 95, 124, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 32, 32, 32, 32, 32, 10},
	},
}

func TestAscii_Output(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Ascii_Output(tt.output, tt.input, tt.args)
			got, _ := os.ReadFile(tt.output)
			if len(got) != len(tt.want) {
				t.Errorf("got %v but wanted %v", string(got), string(tt.want))
			}
			for i := 0; i < len(got); i++ {
				if got[i] != tt.want[i] {
					t.Errorf("got %v but wanted %v", string(got), string(tt.want))
				}
			}
		})
	}
}
