package output

import "testing"


func TestNonPrintable(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonPrintable(tt.args.str); got != tt.want {
				t.Errorf("NonPrintable() = %v, want %v", got, tt.want)
			}
		})
	}
}
