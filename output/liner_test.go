package output

import (
	"reflect"
	"testing"
)

func TestProcessLine(t *testing.T) {
	tests := []struct {
		name string
		line string
		want []int
	}{
		{"hello", "hello", []int{650, 651, 652, 653, 654, 655, 656, 657, 623, 624, 625, 626, 627, 628, 629, 630, 686, 687, 688, 689, 690, 691, 692, 693, 686, 687, 688, 689, 690, 691, 692, 693, 713, 714, 715, 716, 717, 718, 719, 720}},
		{"!", "!",[]int{11 ,12 ,13 ,14 ,15 ,16 ,17 ,18},},
		{"#", "#",[]int{29 ,30 ,31 ,32 ,33 ,34 ,35 ,36},},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProcessLine(tt.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcessLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFndLine(t *testing.T) {
	tests := []struct {
		name string
		args rune
		want []int
	}{
		{"h",'h',[]int{650 ,651 ,652 ,653 ,654 ,655 ,656 ,657},},
		{"!",'!',[]int{11 ,12 ,13 ,14 ,15 ,16 ,17 ,18},},
		{"#",'#',[]int{29 ,30 ,31 ,32 ,33 ,34 ,35 ,36},},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FndLine(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FndLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
