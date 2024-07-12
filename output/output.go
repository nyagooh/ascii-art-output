package output

import (
	"fmt"
)

/*
This function takes a banner file,
opens and reads from it. It also validates if
the passed string input if a printable string.
*/
func ProcessFile(banner string, input string) string {
	var str, filename, printable string
	filename = fmt.Sprintf("%s%s", banner, ".txt")
	printable = NonPrintable(input)
	result := ProcessLine(printable)
	result2 := ProcessString(result, filename)
	str = PrintStrings(result2)

	return str
}
