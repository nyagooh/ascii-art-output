package functions

import "testing"

func TestProcessFile(t *testing.T) {
	tests := []struct {
		name string
		args string
		banner string
		want string
	}{
		{"processfile"},
		{"hello"},
		{"standard"},
		{
			 `
			_              _   _          
		   | |            | | | |         
		   | |__     ___  | | | |   ___   
		   |  _ \   / _ \ | | | |  / _ \  
		   | | | | |  __/ | | | | | (_) | 
		   |_| |_|  \___| |_| |_|  \___/  
										  
										 
		   `
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessFile(tt.banner, tt.args); got != tt.want {
				t.Errorf("ProcessFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
