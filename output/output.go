package output

import (
	"fmt"
	"strings"
)

/*
the function validates the passed string input and splits
by the "\n".
*/
func ProcessFile(banner string, input string) string {
	var str, filename, printable, betterStr string
	filename = fmt.Sprintf("%s%s", banner, ".txt")
	printable = NonPrintable(input)
	if strings.Contains(printable, "\\t") {
		betterStr = strings.ReplaceAll(printable, "\\t", "    ")
	}
	lines := strings.Split(betterStr, "\\n")
	for _, line := range lines {
		if line != "" {
			result := ProcessLine(line)
			result2 := ProcessString(result, filename)
			str = PrintStrings(result2)
		} else {
			fmt.Println()
		}
	}
	return str
}
