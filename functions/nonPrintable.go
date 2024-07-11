package functions

import (
	"fmt"
	// "os"
	"strings"
)

func NonPrintable(str string) string {
	for _, ch := range str {
		if !(ch >= 32 && ch <= 126) {
			fmt.Println("An unprinttable character was detected in your string")
		}
	}
	if strings.Contains(str, "\\t") {
		str = strings.ReplaceAll(str, "\\t", "    ")
	}
	if strings.Contains(str, "\\b") {
		str = strings.ReplaceAll(str, "\\b", " ")
	}
	return str
}
