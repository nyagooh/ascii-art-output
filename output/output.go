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
	var str, filename, printable string
	filename = fmt.Sprintf("%s%s", banner, ".txt")
	printable = NonPrintable(input)
	lines := strings.Split(printable, "\\n")
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
