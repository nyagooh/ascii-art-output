package output

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func Ascii_Output(output, input string, args []string) {
	if !strings.HasSuffix(output, ".txt") {
		log.Fatal("wrong file extension: file must have a '.txt' extension")
	}
	if output == "shadow.txt" || output == "standard.txt" || output == "thinkertoy.txt" {
		log.Fatal("The name of the text file is the same as the banner file name. Please use a different file name.")
	}
	if strings.Count(output, ".") > 1 {
		log.Fatal("Usage: filename.txt")
	}
	var banner string
	if strings.HasSuffix(args[1], ".txt") {
		new := args[:(len(args[1]) - 4)]
		banner = strings.ToLower(strings.Join(new, ""))
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
