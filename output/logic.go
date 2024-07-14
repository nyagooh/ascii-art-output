package output

import (
	"fmt"
	"os"
	"strings"
)

func Ascii_Output(output, input string, args []string) {
	if !strings.HasSuffix(output, ".txt") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if output == "shadow.txt" || output == "standard.txt" || output == "thinkertoy.txt" {
		fmt.Println("The name of the text file is the same as the banner file name. Please use a different file name.")
		os.Exit(0)
	}
	if strings.Count(output, ".") > 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	var banner string
	if strings.HasSuffix(args[1], ".txt") {
		new := strings.TrimSuffix(args[1], ".txt")
		banner = strings.ToLower(new)
	} else {
		banner = strings.ToLower(args[1])
	}
	file, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	lines := strings.Split(input, "\\n")
	for _, line := range lines {
		if line != "" {
			result := ProcessFile(banner, line)
			WriteinFile(file, result)
		} else {
			newline(file)
		}
	}
}

func newline(file *os.File) {
	_, err := file.WriteString("\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func WriteinFile(file *os.File, argument string) {
	_, err := file.WriteString(argument)
	if err != nil {
		fmt.Println(err)
		return
	}
}
