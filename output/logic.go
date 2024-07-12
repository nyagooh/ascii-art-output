package output

import (
	"log"
	"os"
	"strings"
)

func Ascii_Output(output, banner, input string, args []string) {
	if strings.HasSuffix(args[1], ".txt") {
		new := strings.TrimRight(args[1], ".txt")
		banner = strings.ToLower(new)
	} else {
		banner = strings.ToLower(args[1])
	}
	input = args[0]
	if !strings.HasPrefix(os.Args[1], "--output=") {
		log.Fatal("wrong flag format:flag must have '--'")
	}
	output = strings.TrimPrefix(output, "--output=")
	if !strings.HasSuffix(output, ".txt") {
		log.Fatal("wrong file extension: file must have a '.txt' extension")
	}
	if len(output) < 4 {
		log.Fatal("Error: The .txt file size should be greater than 4")
	}
	if output == "shadow.txt" || output == "standard.txt" || output == "thinkertoy.txt" {
		log.Fatal("The name of the text file is the same as the banner file name. Please use a different file name.")
	}
	if strings.Count(output, ".") > 1 {
		log.Fatal("Usage: filename.txt")
	}
	result := ProcessFile(banner, input)
	os.WriteFile(output, []byte(result), 0o644)
}
